package httperror

type HttpError struct {
	Message   string
	Status    string
	ErrorType ErrorType
}

type ErrorType string

const (
	NOT_FOUND     ErrorType = "NOT_FOUND"
	BAD_REQUEST   ErrorType = "BAD_REQUEST"
	ACCESS_DENIED ErrorType = "ACCESS_DENIED"
	INTERNAL                = "INTERNAL"
)

func (h *HttpError) Error() string {
	return h.Message
}

func (h *HttpError) GetStatusCode() int {
	switch h.ErrorType {
	case NOT_FOUND:
		return 404
	case BAD_REQUEST:
		return 400
	case ACCESS_DENIED:
		return 403
	default:
		return 500
	}
}

func NewFailed(m string, e ErrorType) error {
	return &HttpError{
		Message:   m,
		Status:    "FAILED",
		ErrorType: e,
	}
}

func NewInternal(err error) HttpError {
	return HttpError{
		Message:   err.Error(),
		Status:    "ERROR",
		ErrorType: INTERNAL,
	}
}
