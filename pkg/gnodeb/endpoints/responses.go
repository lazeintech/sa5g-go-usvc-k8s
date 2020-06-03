package endpoints

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	_ httptransport.Headerer = (*PreambleResponse)(nil)

	_ httptransport.StatusCoder = (*PreambleResponse)(nil)
)

// PreambleResponse collects the response values for the Sum method.
type PreambleResponse struct {
	Rs  int64 `json:"rs"`
	Err error `json:"err"`
}

func (r PreambleResponse) StatusCode() int {
	return http.StatusOK // TBA
}

func (r PreambleResponse) Headers() http.Header {
	return http.Header{}
}
