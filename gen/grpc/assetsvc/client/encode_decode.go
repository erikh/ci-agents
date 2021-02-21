// Code generated by goa v3.2.6, DO NOT EDIT.
//
// assetsvc gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/tinyci/ci-agents/design

package client

import (
	"context"

	assetsvc "github.com/tinyci/ci-agents/gen/assetsvc"
	assetsvcpb "github.com/tinyci/ci-agents/gen/grpc/assetsvc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildPutLogFunc builds the remote method to invoke for "assetsvc" service
// "putLog" endpoint.
func BuildPutLogFunc(grpccli assetsvcpb.AssetsvcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.PutLog(ctx, opts...)
		}
		return grpccli.PutLog(ctx, opts...)
	}
}

// DecodePutLogResponse decodes responses from the assetsvc putLog endpoint.
func DecodePutLogResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	return &PutLogClientStream{
		stream: v.(assetsvcpb.Assetsvc_PutLogClient),
	}, nil
}

// BuildGetLogFunc builds the remote method to invoke for "assetsvc" service
// "getLog" endpoint.
func BuildGetLogFunc(grpccli assetsvcpb.AssetsvcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetLog(ctx, reqpb.(*assetsvcpb.GetLogRequest), opts...)
		}
		return grpccli.GetLog(ctx, &assetsvcpb.GetLogRequest{}, opts...)
	}
}

// EncodeGetLogRequest encodes requests sent to assetsvc getLog endpoint.
func EncodeGetLogRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*assetsvc.GetLogPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("assetsvc", "getLog", "*assetsvc.GetLogPayload", v)
	}
	return NewGetLogRequest(payload), nil
}

// DecodeGetLogResponse decodes responses from the assetsvc getLog endpoint.
func DecodeGetLogResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	return &GetLogClientStream{
		stream: v.(assetsvcpb.Assetsvc_GetLogClient),
	}, nil
}
