package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	transport "github.com/erikh/go-transport"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	opentracing "github.com/opentracing/opentracing-go"
	apiSess "github.com/tinyci/ci-agents/api/sessions"
	"github.com/tinyci/ci-agents/clients/github"
	"github.com/tinyci/ci-agents/clients/log"
	"github.com/tinyci/ci-agents/config"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/model"
	"github.com/tinyci/ci-agents/types"
	"github.com/tinyci/ci-agents/utils"
	"golang.org/x/net/websocket"
)

// SessionUsername is the name of the session key that contains our username value.
const SessionUsername = "username"

// AllowOrigin changes the scope of CORS requests. By default they are
// insecure; this is intended to be overrided by commands as they set up this
// library along with their other standard init operations.
var AllowOrigin = "*"

// ErrRedirect indicates that the error intends to redirect the user to the proper spot.
var ErrRedirect = errors.New("redirection")

var (
	errNoCapability  = errors.New("no capability to perform desired operation")
	errInvalidCookie = errors.New("cookie was invalid")

	routeTransformer = regexp.MustCompile(`(?:{([^}]+)})+`)
)

// HandlerConfig provides an interface to managing the HandlerConfig.
type HandlerConfig interface {
	SetRoutes(*H)
	DBConfigure(*H) *errors.Error
	Configure(Routes) *errors.Error
	CustomInit(*H) *errors.Error
	Validate(*H) *errors.Error
}

// H is a series of HTTP handlers for the UI service
type H struct {
	Config HandlerConfig `yaml:"-"`
	Routes Routes        `yaml:"-"`

	config.UserConfig `yaml:",inline"`
	config.Service
}

// GetUser retrieves the user based on information in the gin context.
func (h *H) GetUser(ctx *gin.Context) (*model.User, *errors.Error) {
	client, err := h.GetClient(ctx)
	if err != nil {
		return nil, err
	}

	var name string
	sess := h.Session(ctx)

	username := sess.Get(SessionUsername)

	// FIXME clean up this spaghetti. Too much branching and ultimately, we're
	// looking to get this into the `name` or `u` variables and that isn't very
	// obvious.  Make this a function. -erikh
	var u *model.User

	if username == nil {
		if token := ctx.Request.Header.Get("Authorization"); token != "" {
			token := ctx.Request.Header.Get("Authorization")
			if token != "" {
				u, err = h.Clients.Data.ValidateToken(ctx, token)
				if err != nil {
					return nil, err
				}
			}
		} else {
			var err *errors.Error
			name, err = client.MyLogin(ctx)
			if err != nil {
				return nil, errors.New(err)
			}

			sess.Set(SessionUsername, name)
			if err := sess.Save(); err != nil {
				return nil, errors.New(err)
			}
		}
	} else {
		var ok bool
		name, ok = username.(string)
		if !ok {
			return nil, errors.ErrInvalidAuth
		}
	}

	if u == nil && name != "" {
		u, err = h.Clients.Data.GetUser(ctx, name)
		if err != nil {
			return nil, err
		}
	}

	return u, nil
}

// GetClient returns a github client that works with the credentials in the given context.
func (h *H) GetClient(ctx *gin.Context) (github.Client, *errors.Error) {
	user, err := h.GetGithub(ctx)
	if err != nil {
		return nil, err
	}

	token := &types.OAuthToken{}

	if err := utils.JSONIO(user.Token, token); err != nil {
		return nil, err
	}

	return h.GithubClient(token), nil
}

// OAuthRedirect redirects the user to the OAuth redirection URL.
func (h *H) OAuthRedirect(ctx *gin.Context, scopes []string) *errors.Error {
	url, err := h.Clients.Auth.GetOAuthURL(ctx, scopes)
	if err != nil {
		return err
	}

	ctx.Redirect(302, url)
	return nil
}

// GithubClient is a wrapper for config.GithubClient.
func (h *H) GithubClient(token *types.OAuthToken) github.Client {
	return h.OAuth.GithubClient(token)
}

// CreateClients creates the clients to be used based on configuration values.
func (h *H) CreateClients() *errors.Error {
	var err *errors.Error
	h.Clients, err = h.ClientConfig.CreateClients(h.UserConfig, h.Name)

	return err
}

// CreateTransport creates a transport with optional certification information.
func (h *H) CreateTransport() (*transport.HTTP, *errors.Error) {
	var cert *transport.Cert
	if !h.NoTLSServer {
		var err error
		cert, err = h.TLS.Load()
		if err != nil {
			return nil, errors.New(err)
		}
	}

	t, err := transport.NewHTTP(cert)
	if err != nil {
		return nil, errors.New(err)
	}

	return t, nil
}

// Boot boots the service. Closing the channel returned will shutdown the
// service. At shutdown time, this routine will close the finished channel when
// it is finished shutting everything down, so the program can safely
// terminate.
func Boot(t *transport.HTTP, handler *H, finished chan struct{}) (chan struct{}, *errors.Error) {
	handler.Formats = strfmt.NewFormats()

	if err := handler.Init(); err != nil {
		return nil, err
	}

	var (
		closer io.Closer
		err    *errors.Error
	)

	if handler.EnableTracing {
		closer, err = handler.createGlobalTracer()
		if err != nil {
			return nil, err
		}
	}

	r, err := handler.CreateRouter()
	if err != nil {
		return nil, err
	}

	if t == nil {
		var err *errors.Error
		t, err = handler.CreateTransport()
		if err != nil {
			return nil, err
		}
	}

	var sErr error
	s, l, sErr := t.Server(fmt.Sprintf(":%d", handler.Port), r)
	if sErr != nil {
		return nil, errors.New(sErr)
	}

	s.IdleTimeout = 30 * time.Second
	doneChan := make(chan struct{})

	go func() {
		<-doneChan
		s.Close()
		l.Close()
		handler.Clients.CloseClients()
		if handler.EnableTracing {
			closer.Close()
		}
		close(finished)
	}()

	go func() {
		if err := s.Serve(l); err != nil {
			handler.Clients.Log.Error(context.Background(), err)
		}
	}()
	return doneChan, nil
}

// NewTracingSpan creates a new tracepoint span for opentracing instrumentation.
func (h *H) NewTracingSpan(ctx *gin.Context, operation string) opentracing.Span {
	span, ctx2 := opentracing.StartSpanFromContext(ctx.Request.Context(), operation)
	ctx.Request = ctx.Request.WithContext(ctx2)
	ctx.Next()

	return span
}

func (h *H) createGlobalTracer() (io.Closer, *errors.Error) {
	return utils.CreateTracer("uisvc")
}

// Init initialize the handler and makes it available for requests.
func (h *H) Init() *errors.Error {
	if err := h.Config.Validate(h); err != nil {
		return err
	}

	if err := h.Config.CustomInit(h); err != nil {
		return err
	}

	if err := h.CreateClients(); err != nil {
		return err
	}

	h.Config.SetRoutes(h)

	if err := h.Config.Configure(h.Routes); err != nil {
		return err
	}

	if err := h.Auth.Validate(h.UseSessions); err != nil {
		return err
	}

	return h.dbConnect()
}

func (h *H) dbConnect() *errors.Error {
	if h.UseDB {
		var err *errors.Error
		h.Model, err = model.New(h.DSN)
		if err != nil {
			return err
		}
		return h.Config.DBConfigure(h)
	}

	return nil
}

// CORS primes OPTIONS and normal requests with the appropriate headers and
// also acts like a normal http.Handler so it can be used that way.
func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", AllowOrigin)
	ctx.Header("Access-Control-Allow-Methods", "PUT,GET,POST,DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")
}

// GetGithub gets the github user from the session and loads it.
func (h *H) GetGithub(ctx *gin.Context) (u *model.User, outErr *errors.Error) {
	sess := sessions.Default(ctx)

	defer func() {
		if outErr != nil {
			sess.Clear()
		}
	}()

	uname, ok := sess.Get(SessionUsername).(string)
	if ok && strings.TrimSpace(uname) != "" {
		// no error, we're already logged in
		return h.Clients.Data.GetUser(ctx, uname)
	}

	token := ctx.Request.Header.Get("Authorization")
	if token != "" {
		return h.Clients.Data.ValidateToken(ctx, token)
	}

	return nil, errInvalidCookie
}

func (h *H) authed(gatewayFunc func(*H, *gin.Context, HandlerFunc) *errors.Error, cap model.Capability, scope string) func(h *H, ctx *gin.Context, processor HandlerFunc) *errors.Error {
	return func(h *H, ctx *gin.Context, processor HandlerFunc) *errors.Error {
		var (
			u            *model.User
			err          *errors.Error
			span         opentracing.Span
			spanFinished bool
		)

		if h.EnableTracing {
			span = h.NewTracingSpan(ctx, "auth exchange")
			defer func() {
				if !spanFinished {
					span.Finish()
				}
			}()
		}

		token := ctx.Request.Header.Get("Authorization")
		if token != "" {
			if h.EnableTracing {
				span.LogKV("authorization", "token")
			}
			u, err = h.Clients.Data.ValidateToken(ctx, token)
			if err != nil {
				return err
			}
		} else {
			if h.EnableTracing {
				span.LogKV("authorization", "github")
			}

			u, err = h.GetGithub(ctx)
			if err != nil {
				return err
			}
		}

		if cap != "" {
			if h.EnableTracing {
				span.LogKV("event", "capability check")
			}

			res, err := h.Clients.Data.HasCapability(ctx, u, cap)
			if err != nil {
				return err
			}

			if !res {
				return errNoCapability
			}
		}

		if scope != "" && !u.Token.Can(scope) {
			return errors.New("cannot perform operation with current oauth scopes; must upgrade")
		}

		if h.EnableTracing {
			spanFinished = true
			span.Finish()
		}

		return gatewayFunc(h, ctx, processor)
	}
}

func (h *H) inWebsocket(key string, paramHandler func(*H, *gin.Context) *errors.Error, handler WebsocketFunc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if h.EnableTracing {
			span := h.NewTracingSpan(ctx, key)
			defer span.Finish()
		}

		outerHandler := func(conn *websocket.Conn) {
			pCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()
			defer conn.Close()
			if err := paramHandler(h, ctx); err != nil {
				// FIXME log
				return
			}

			if err := handler(pCtx, h, ctx, conn); err != nil {
				// FIXME log
				return
			}
		}

		if h.EnableTracing {
			span2 := h.NewTracingSpan(ctx, "websocket communication")
			defer span2.Finish()
		}
		websocket.Handler(outerHandler).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// TransformSwaggerRoute merely translates url params from {thisformat} to :thisformat
func TransformSwaggerRoute(route string) string {
	return routeTransformer.ReplaceAllStringFunc(route, func(input string) string {
		return string(routeTransformer.ExpandString([]byte{}, ":$1", input, routeTransformer.FindStringSubmatchIndex(input)))
	})
}

func (h *H) configureSessions(r *gin.Engine) *errors.Error {
	sessdb := apiSess.New(h.Clients.Data, nil, h.Auth.ParsedSessionCryptKey())
	r.Use(sessions.Sessions(config.SessionKey, sessdb))

	return nil
}

func (h *H) configureRestHandler(r *gin.Engine, key string, route *Route, optionsRoutes map[string]struct{}) {
	var dispatchFunc func(string, ...gin.HandlerFunc) gin.IRoutes

	switch route.Method {
	case "GET":
		dispatchFunc = r.GET
	case "POST":
		dispatchFunc = r.POST
	case "DELETE":
		dispatchFunc = r.DELETE
	case "PATCH":
		dispatchFunc = r.PATCH
	case "PUT":
		dispatchFunc = r.PUT
	case "OPTIONS":
		dispatchFunc = r.OPTIONS
	case "HEAD":
		dispatchFunc = r.HEAD
	}

	var handler func(*H, *gin.Context, HandlerFunc) *errors.Error = route.Handler

	if route.UseAuth {
		handler = h.authed(handler, route.Capability, route.TokenScope)
	}

	if route.UseCORS {
		if _, ok := optionsRoutes[key]; !ok {
			r.OPTIONS(key, CORS)
			optionsRoutes[key] = struct{}{}
		}
	}
	dispatchFunc(key, h.wrapHandler(key, handler, route.Processor))
}

// CreateRouter creates a *mux.Router capable of serving the UI server.
func (h *H) CreateRouter() (*gin.Engine, *errors.Error) {
	r := gin.New()
	r.Use(ginhttp.Middleware(opentracing.GlobalTracer()))

	if h.UseSessions {
		if err := h.configureSessions(r); err != nil {
			return nil, err
		}
	}

	optionsRoutes := map[string]struct{}{}

	for key, methodRoutes := range h.Routes {
		for _, route := range methodRoutes {
			if route.WebsocketProcessor != nil {
				r.GET(key, h.inWebsocket(key, route.ParamValidator, route.WebsocketProcessor))
			} else {
				h.configureRestHandler(r, key, route, optionsRoutes)
			}
		}
	}

	return r, nil
}

func (h *H) wrapHandler(key string, handler func(*H, *gin.Context, HandlerFunc) *errors.Error, processor HandlerFunc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if h.EnableTracing {
			span := h.NewTracingSpan(ctx, key)
			defer span.Finish()
		}

		if err := handler(h, ctx, processor); err != nil {
			if err.Contains(ErrRedirect) {
				return
			}

			h.WriteError(ctx, err)
		}
	}
}

// WriteError standardizes the writing of error states for easier typing. It is
// not intended to be used to write specific statuses, only 500 errors with JSON output.
// If UseSessions is on, it will populate the errors session store.
func (h *H) WriteError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(500, err)
}

// Session returns the current user session.
func (h *H) Session(ctx *gin.Context) sessions.Session {
	return sessions.Default(ctx)
}

// LogError logs an HTTP error to the client.
func (h *H) LogError(err error, ctx *gin.Context, code int) {
	logger := h.Clients.Log.WithRequest(ctx.Request).WithFields(log.Fields{"code": fmt.Sprintf("%v", code)})
	user, gitErr := h.GetGithub(ctx)
	if gitErr == nil {
		logger = logger.WithUser(user)
	}

	content, jsonErr := json.Marshal(ctx.Params)
	if jsonErr != nil {
		logger.Error(context.Background(), errors.New(jsonErr).Wrap("encoding params for log message"))
	}

	var doLog bool

	switch err := err.(type) {
	case *errors.Error:
		if err.GetLog() {
			doLog = true
		}
	default:
		doLog = true
	}

	if doLog {
		logger.WithFields(log.Fields{"params": string(content)}).Error(context.Background(), err)
	}
}
