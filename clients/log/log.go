// Package log is a singleton package for logging either remotely or
// locally. If provided a host, it will connect to a service utilizing the
// syslogsvc protocol for tinyCI, and will alloow it to send log transmissions
// there.
package log

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	transport "github.com/erikh/go-transport"
	"github.com/sirupsen/logrus"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/gen/grpc/logsvc/client"
	"github.com/tinyci/ci-agents/gen/logsvc"
	"github.com/tinyci/ci-agents/model"
	"github.com/tinyci/ci-agents/utils"
	"google.golang.org/grpc"
)

// Client is a small wrapper around the GRPC client that includes tracing data.
type Client struct {
	conn   *grpc.ClientConn
	closer io.Closer
	closed bool
}

const (
	// LevelDebug is the debug loglevel
	LevelDebug = "debug"
	// LevelError is the error loglevel
	LevelError = "error"
	// LevelInfo is the info loglevel
	LevelInfo = "info"
)

// Close closes the client's tracing functionality
func (c *Client) Close() *errors.Error {
	remoteMutex.Lock()
	defer remoteMutex.Unlock()

	if !c.closed && c.closer != nil {
		c.closed = true
		return errors.New(c.closer.Close())
	}

	return nil
}

func (c *Client) goaClient() *client.Client {
	return client.NewClient(c.conn)
}

var (
	// RemoteClient is the swagger-based syslogsvc client.
	RemoteClient *Client
	// remoteMutex is the mutex used to control setting it
	remoteMutex sync.Mutex
)

// Fields is just shorthand for some protobuf stuff.
type Fields map[string]string

// ToLogrus converts Fields to logrus.Fields
func (f Fields) ToLogrus() logrus.Fields {
	ret := logrus.Fields{}

	for key, value := range f {
		ret[key] = value
	}

	return ret
}

// ConfigureRemote configures the remote endpoint with a provided URL.
func ConfigureRemote(addr string, cert *transport.Cert, trace bool) *errors.Error {
	remoteMutex.Lock()
	defer remoteMutex.Unlock()

	var (
		closer  io.Closer
		options []grpc.DialOption
		eErr    *errors.Error
	)

	if trace {
		closer, options, eErr = utils.SetUpGRPCTracing("log")
		if eErr != nil {
			return eErr
		}
	}

	c, err := transport.GRPCDial(cert, addr, options...)
	if err != nil {
		return errors.New(err)
	}

	RemoteClient = &Client{conn: c, closer: closer}
	return nil
}

// SubLogger is a handle to cached parameters for logging.
type SubLogger struct {
	Service string
	Fields  Fields
}

// New creates a new SubLogger which can be primed with cached values for each log entry.
func New() *SubLogger {
	return &SubLogger{Service: "n/a", Fields: Fields{}}
}

// NewWithData returns a SubLogger already primed with cached data.
func NewWithData(svc string, params Fields) *SubLogger {
	if params == nil {
		params = Fields{}
	}

	params["service"] = svc
	return &SubLogger{svc, params}
}

// WithService is a SubLogger version of package-level WithService. They call the same code.
func (sub *SubLogger) WithService(svc string) *SubLogger {
	sub2 := *sub
	sub2.Service = svc
	sub2.Fields = Fields{}

	if sub.Fields != nil {
		params := sub.Fields

		for k, v := range params {
			sub2.Fields[k] = v
		}
	}

	sub2.Fields["service"] = svc
	return &sub2
}

// WithFields is a SubLogger version of package-level WithFields. They call the same code.
func (sub *SubLogger) WithFields(params Fields) *SubLogger {
	sub2 := *sub

	sub2.Fields = Fields{}

	for k, v := range sub.Fields {
		sub2.Fields[k] = v
	}

	for k, v := range params {
		sub2.Fields[k] = v
	}

	return &sub2
}

// WithRequest is a wrapper for WithFields() that handles *http.Request data.
func (sub *SubLogger) WithRequest(req *http.Request) *SubLogger {
	raddr := req.Header.Get("X-Forwarded-For")
	if raddr == "" {
		raddr = strings.Split(req.RemoteAddr, ":")[0]
	} else {
		raddr = strings.TrimSpace(strings.SplitN(raddr, ",", 2)[0])
	}

	fm := Fields{
		"remote_addr":    raddr,
		"request_method": req.Method,
		"request_url":    req.URL.String(),
	}

	return sub.WithFields(fm)
}

// WithUser includes user information
func (sub *SubLogger) WithUser(user *model.User) *SubLogger {
	fm := Fields{
		"username": user.Username,
		"user_id":  fmt.Sprintf("%v", user.ID),
	}

	return sub.WithFields(fm)
}

func (sub *SubLogger) makeMsg(level, msg string, values []interface{}) *logsvc.PutPayload {
	if values != nil {
		msg = fmt.Sprintf(msg, values...)
	}

	return &logsvc.PutPayload{
		At:      time.Now().Unix(),
		Fields:  sub.Fields,
		Message: msg,
		Level:   level,
		Service: sub.Service,
	}
}

// Logf logs a thing with formats!
func (sub *SubLogger) Logf(ctx context.Context, level string, msg string, values []interface{}, localLog func(string, ...interface{})) {
	if RemoteClient != nil {
		made := sub.makeMsg(level, msg, values)

		if _, err := RemoteClient.goaClient().Put()(ctx, made); err != nil {
			localLog(err.Error())
			localLog("%v", made)
			return
		}
	}

	localLog(msg, values...)
}

// Log logs a thing
func (sub *SubLogger) Log(ctx context.Context, level string, msg interface{}, localLog func(...interface{})) {
	if RemoteClient != nil {
		var made *logsvc.PutPayload

		switch msg := msg.(type) {
		case string:
			made = sub.makeMsg(level, msg, nil)
		case fmt.Stringer:
			made = sub.makeMsg(level, msg.String(), nil)
		case errors.Error:
			made = sub.makeMsg(level, msg.Error(), nil)
		case logsvc.PutPayload:
			made = sub.makeMsg(level, msg.Message, nil)
		default:
			localLog(fmt.Sprintf("attempted to log an invalid value (%T); coerce to fmt.Stringer or error first", msg))
			return
		}

		if _, err := RemoteClient.goaClient().Put()(ctx, made); err != nil {
			localLog(err)
			localLog(made)
			return
		}
	}

	switch msg := msg.(type) {
	case errors.Error:
		if msg.Log {
			localLog(msg)
		}
	default:
		localLog(msg)
	}
}

// Info prints an info message
func (sub *SubLogger) Info(ctx context.Context, msg interface{}) {
	sub.Log(ctx, LevelInfo, msg, logrus.WithFields(sub.Fields.ToLogrus()).Info)
}

// Infof is the format-capable version of Info
func (sub *SubLogger) Infof(ctx context.Context, msg string, values ...interface{}) {
	sub.Logf(ctx, LevelInfo, msg, values, logrus.WithFields(sub.Fields.ToLogrus()).Infof)
}

// Error prints an error message
func (sub *SubLogger) Error(ctx context.Context, msg interface{}) {
	sub.Log(ctx, LevelError, msg, logrus.WithFields(sub.Fields.ToLogrus()).Error)
}

// Errorf is the format-capable version of Error
func (sub *SubLogger) Errorf(ctx context.Context, msg string, values ...interface{}) {
	sub.Logf(ctx, LevelError, msg, values, logrus.WithFields(sub.Fields.ToLogrus()).Errorf)
}

// Debug prints a debug message
func (sub *SubLogger) Debug(ctx context.Context, msg interface{}) {
	sub.Log(ctx, LevelDebug, msg, logrus.WithFields(sub.Fields.ToLogrus()).Debug)
}

// Debugf is the format-capable version of Debug
func (sub *SubLogger) Debugf(ctx context.Context, msg string, values ...interface{}) {
	sub.Logf(ctx, LevelDebug, msg, values, logrus.WithFields(sub.Fields.ToLogrus()).Debugf)
}
