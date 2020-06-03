package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"

	"github.com/miki-tnt/sa5g-go-usvc-k8s/pkg/addsvc/service"
)

// Endpoints collects all of the endpoints that compose the addsvc service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	PreambleEndpoint endpoint.Endpoint `json:""`
}

// New return a new instance of the endpoint that wraps the provided service.
func New(svc service.AddsvcService, logger log.Logger, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer) (ep Endpoints) {
	var preambleEndpoint endpoint.Endpoint
	{
		method := "sum"
		preambleEndpoint = MakePreambleEndpoint(svc)
		preambleEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))(preambleEndpoint)
		preambleEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(preambleEndpoint)
		preambleEndpoint = opentracing.TraceServer(otTracer, method)(preambleEndpoint)
		preambleEndpoint = zipkin.TraceEndpoint(zipkinTracer, method)(preambleEndpoint)
		preambleEndpoint = LoggingMiddleware(log.With(logger, "method", method))(preambleEndpoint)
		ep.PreambleEndpoint = preambleEndpoint
	}

	return ep
}

// MakePreambleEndpoint returns an endpoint that invokes Preamble on the service.
// Primarily useful in a server.
func MakePreambleEndpoint(svc service.AddsvcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PreambleRequest)
		if err := req.validate(); err != nil {
			return PreambleResponse{}, err
		}
		rs, err := svc.Preamble(ctx, req.Msg)
		return PreambleResponse{Rs: rs}, err
	}
}

// Preamble implements the service interface, so Endpoints may be used as a service.
// This is primarily useful in the context of a client library.
func (e Endpoints) Preamble(ctx context.Context, msg int64) (rs int64, err error) {
	resp, err := e.PreambleEndpoint(ctx, PreambleRequest{Msg: msg})
	if err != nil {
		return
	}
	response := resp.(PreambleResponse)
	return response.Rs, nil
}
