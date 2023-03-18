package error_utils

import (
	"encoding/json"
	"net/http"
)

type AppErr interface {
	Message() string
	Status() int
	Error() string
}

type appErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *appErr) Error() string {
	return e.ErrError
}

func (e *appErr) Message() string {
	return e.ErrMessage
}

func (e *appErr) Status() int {
	return e.ErrStatus
}

func AppNotFoundError(message string) AppErr {
	return &appErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func AppBadRequestError(message string) AppErr {
	return &appErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func AppUnprocessibleEntityError(message string) AppErr {
	return &appErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "invalid_request",
	}
}

func AppInternalServerError(message string) AppErr {
	return &appErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "server_error",
	}
}

func AppApiBytesError (body []byte) (AppErr, error) {
	var result appErr
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
