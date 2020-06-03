package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(PreamblesvcService) PreamblesvcService

// Service describes a service that adds things together
// Implement yor service methods methods.
// e.x: Foo(ctx context.Context, s string)(rs string, err error)
type PreamblesvcService interface {
	Preamble(ctx context.Context, msg int64) (rs int64, err error)
}

// the concrete implementation of service interface
type stubPreamblesvcService struct {
	logger log.Logger `json:"logger"`
}

// New return a new instance of the service.
// If you want to add service middleware this is the place to put them.
func New(logger log.Logger) (s PreamblesvcService) {
	var svc PreamblesvcService
	{
		svc = &stubPreamblesvcService{logger: logger}
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

// Implement the business logic of Preamble
func (ad *stubPreamblesvcService) Preamble(ctx context.Context, msg int64) (rs int64, err error) {
	return msg, err
}
