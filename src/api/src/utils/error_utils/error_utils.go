package error_utils

import (
	"encoding/json"
	"net/http"
)

type TodoErr interface {
	Message() string
	Status() int
	Error() string
}

type todoErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *todoErr) Error() string {
	return e.ErrError
}

func (e *todoErr) Message() string {
	return e.ErrMessage
}

func (e *todoErr) Status() int {
	return e.ErrStatus
}

func NewNotFoundError(message string) TodoErr {
	return &todoErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewBadRequestError(message string) TodoErr {
	return &todoErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}
func NewUnprocessibleEntityError(message string) TodoErr {
	return &todoErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "invalid_request",
	}
}

func NewApiErrFromBytes(body []byte) (TodoErr, error) {
	var result todoErr
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func NewInternalServerError(message string) TodoErr {
	return &todoErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "server_error",
	}
}