// Code generated by goa v3.2.6, DO NOT EDIT.
//
// logsvc gRPC server
//
// Command:
// $ goa gen github.com/tinyci/ci-agents/design

package server

import (
	"context"

	logsvcpb "github.com/tinyci/ci-agents/gen/grpc/logsvc/pb"
	logsvc "github.com/tinyci/ci-agents/gen/logsvc"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the logsvcpb.LogsvcServer interface.
type Server struct {
	PutH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the logsvc service endpoints.
func New(e *logsvc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		PutH: NewPutHandler(e.Put, uh),
	}
}

// NewPutHandler creates a gRPC handler which serves the "logsvc" service "put"
// endpoint.
func NewPutHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodePutRequest, EncodePutResponse)
	}
	return h
}

// Put implements the "Put" method in logsvcpb.LogsvcServer interface.
func (s *Server) Put(ctx context.Context, message *logsvcpb.PutRequest) (*logsvcpb.PutResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "put")
	ctx = context.WithValue(ctx, goa.ServiceKey, "logsvc")
	resp, err := s.PutH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*logsvcpb.PutResponse), nil
}
