package errors

import (
	error_helper "github.com/a666hn/go-libs/error/helper"
	"net/http"
)

type HTTPError struct {
	CommonError
	HttpStatusNumber int    `json:"-"`
	HttpStatusName   string `json:"type"`
}

func (h *HTTPError) Error() string {
	return h.ClientMessage
}

func (err *CommonError) GetHttpStatus() int {
	if errorDictionaries.httpCodes[err.ErrorCode] == 0 {
		return http.StatusInternalServerError
	}
	return errorDictionaries.httpCodes[err.ErrorCode]
}

func (err *CommonError) ToHttpError() HTTPError {
	s := err.GetHttpStatus()
	return HTTPError{
		CommonError:      *err,
		HttpStatusNumber: s,
		HttpStatusName:   error_helper.GetHttpStatusName(s),
	}
}
