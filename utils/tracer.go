package utils

import (
	"io"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
)

func getConfig() (*config.Configuration, error) {
	// FIXME Taken from jaeger/opentracing examples; needs tunables.
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	cfg.Sampler = &config.SamplerConfig{
		Type:  "const",
		Param: 1,
	}

	return cfg, nil
}

// CreateTracer creates an opentracing-compatible jaegertracing client.
func CreateTracer(serviceName string) (io.Closer, error) {
	cfg, eErr := getConfig()
	if eErr != nil {
		return nil, eErr
	}

	closer, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		return nil, err
	}

	return closer, nil
}

// SetUpGRPCTracing configures grpc dial functions for tracing.
func SetUpGRPCTracing(client string) (io.Closer, []grpc.DialOption, error) {
	cfg, eErr := getConfig()
	if eErr != nil {
		return nil, nil, eErr
	}

	cfg.ServiceName = client
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}

	return closer,
		[]grpc.DialOption{
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
			grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)),
		}, nil
}
