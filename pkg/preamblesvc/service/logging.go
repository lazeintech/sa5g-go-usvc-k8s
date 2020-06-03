package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type loggingMiddleware struct {
	logger log.Logger         `json:""`
	next   PreamblesvcService `json:""`
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PreamblesvcService) PreamblesvcService {
		return loggingMiddleware{level.Info(logger), next}
	}
}

func (lm loggingMiddleware) Preamble(ctx context.Context, msg int64) (rs int64, err error) {
	defer func(begin time.Time) {
		lm.logger.Log("method", "Preable", "msg", msg, "err", err)
	}(time.Now())

	return lm.next.Preamble(ctx, a, b)
}
