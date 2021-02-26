package asset

import (
	"context"
	"io"

	transport "github.com/erikh/go-transport"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/gen/assetsvc"
	"github.com/tinyci/ci-agents/gen/grpc/assetsvc/client"
	"github.com/tinyci/ci-agents/utils"
	"google.golang.org/grpc"
)

// Client is a handle into the asset client.
type Client struct {
	ac     *client.Client
	closer io.Closer
}

// NewClient creates a new *Client for use.
func NewClient(addr string, cert *transport.Cert, trace bool) (*Client, *errors.Error) {
	var (
		closer  io.Closer
		options []grpc.DialOption
		eErr    *errors.Error
	)

	if trace {
		closer, options, eErr = utils.SetUpGRPCTracing("asset")
		if eErr != nil {
			return nil, eErr
		}
	}

	t, err := transport.GRPCDial(cert, addr, options...)
	if err != nil {
		return nil, errors.New(err)
	}
	return &Client{closer: closer, ac: client.NewClient(t)}, nil
}

// Close closes the client's tracing functionality
func (c *Client) Close() *errors.Error {
	if c.closer != nil {
		return errors.New(c.closer.Close())
	}

	return nil
}

// Write writes a log at id with the supplied reader providing the content.
func (c *Client) Write(ctx context.Context, id int64, f io.Reader) *errors.Error {
	l, err := c.ac.PutLog()(ctx, nil)
	if err != nil {
		return errors.New(err)
	}

	log := l.(*client.PutLogClientStream)

	buf := make([]byte, 64)

	for {
		var done bool
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return errors.New(err)
		} else if err == io.EOF {
			done = true
		}

		ls := &assetsvc.PutLogStreamingPayload{
			ID:    id,
			Chunk: buf[:n],
		}

		if err := log.Send(ls); err != nil && err != io.EOF {
			return errors.New(err)
		} else if err == io.EOF {
			done = true
		}

		if done {
			if err := log.Close(); err != nil {
				return errors.New(err)
			}

			return nil
		}
	}
}

func (c *Client) Read(ctx context.Context, id int64, w io.Writer) *errors.Error {
	as, err := c.ac.GetLog()(ctx, &assetsvc.GetLogPayload{ID: id})
	if err != nil {
		return errors.New(err)
	}

	log := as.(*client.GetLogClientStream)

	for {
		chunk, err := log.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return errors.New(err)
		}

		if errs := chunk.Errors; len(errs) > 0 {
			return errors.New(errs[0])
		}

		if _, err := w.Write(chunk.Chunk); err != nil {
			return errors.New(err)
		}
	}
}
