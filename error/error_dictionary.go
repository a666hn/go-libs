package errors

import "net/http"

type ErrorCode uint

const UNKNOWN_ERROR ErrorCode = 0

var errorDictionaries *ErrorDictionaries = &ErrorDictionaries{
	errorCodes: map[ErrorCode]*CommonError{},
	httpCodes:  map[ErrorCode]int{},
}

type ErrorDictionaries struct {
	errorCodes map[ErrorCode]*CommonError
	httpCodes  map[ErrorCode]int
}

func NewErrorDictionariesInstance(
	errorCode map[ErrorCode]*CommonError,
	httpCode map[ErrorCode]int,
) *ErrorDictionaries {
	errorDictionaries = &ErrorDictionaries{
		errorCodes: errorCode,
		httpCodes:  httpCode,
	}

	errorDictionaries.errorCodes[UNKNOWN_ERROR] = &CommonError{
		ClientMessage:         "Unhandled error.",
		ServiceMessage:        "An unhandled error has occured. Please contact the developer.",
		OzzoValidationMessage: nil,
		ErrorCode:             UNKNOWN_ERROR,
	}

	errorDictionaries.httpCodes[UNKNOWN_ERROR] = http.StatusInternalServerError

	return errorDictionaries
}
