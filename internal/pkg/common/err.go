package common

import (
	"net/http"
)

type ErrResponseDto struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int    `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponseDto) IsEmpty() bool {
	return e.Err == nil && e.HTTPStatusCode == 0 &&
		e.StatusText == "" && e.AppCode == 0 &&
		e.ErrorText == ""
}

func (e ErrResponseDto) Render(w http.ResponseWriter, _ *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) ErrResponseDto {
	code := http.StatusBadRequest
	return ErrResponseDto{
		Err:            err,
		HTTPStatusCode: code,
		StatusText:     http.StatusText(code),
		ErrorText:      err.Error(),
	}
}

func ErrInternalServerError(err error) ErrResponseDto {
	code := http.StatusInternalServerError
	return ErrResponseDto{
		Err:            err,
		HTTPStatusCode: code,
		StatusText:     http.StatusText(code),
		ErrorText:      err.Error(),
	}
}

func ErrNotFound() ErrResponseDto {
	code := http.StatusNotFound
	return ErrResponseDto{
		HTTPStatusCode: code,
		StatusText:     http.StatusText(code),
	}
}
