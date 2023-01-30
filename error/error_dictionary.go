package errors

import "net/http"

type ErrorCode uint

const UNKNOWN_ERROR ErrorCode = 0

var eDicts = &ErrorDictionaries{
	errorCodes: make(map[ErrorCode]*CommonError),
	httpCodes:  make(map[ErrorCode]int),
}

type ErrorDictionaries struct {
	errorCodes map[ErrorCode]*CommonError
	httpCodes  map[ErrorCode]int
}

func NewErrorDictionariesInstance(
	errorCode map[ErrorCode]*CommonError,
	httpCode map[ErrorCode]int,
) *ErrorDictionaries {
	eDicts = &ErrorDictionaries{
		errorCodes: errorCode,
		httpCodes:  httpCode,
	}

	eDicts.errorCodes[UNKNOWN_ERROR] = &CommonError{
		ClientMessage:         "Unhandled error.",
		ServiceMessage:        "An unhandled error has occured. Please contact the developer.",
		OzzoValidationMessage: nil,
		ErrorCode:             UNKNOWN_ERROR,
	}

	eDicts.httpCodes[UNKNOWN_ERROR] = http.StatusInternalServerError

	return eDicts
}
