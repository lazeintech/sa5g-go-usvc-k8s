package endpoints

type Request interface {
	validate() error
}

// PreambleRequest collects the request parameters for the Sum method.
type PreambleRequest struct {
	Msg int64 `json:"msg"`
}

func (r PreambleRequest) validate() error {
	return nil // TBA
}
