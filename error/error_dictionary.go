package errors

import "net/http"

type ErrorCode uint

const UNKNOWN_ERROR ErrorCode = 0

type ErrorDictionaries struct {
	errorCode map[ErrorCode]*CommonError
	httpCode  map[ErrorCode]int
}

var errorDictionaries = &ErrorDictionaries{
	errorCode: make(map[ErrorCode]*CommonError),
	httpCode:  make(map[ErrorCode]int),
}

func RegisterErrorDictionaries(
	errorCode map[ErrorCode]*CommonError,
	httpCode map[ErrorCode]int,
) *ErrorDictionaries {
	errorDicts := &ErrorDictionaries{
		errorCode: errorCode,
		httpCode:  httpCode,
	}

	errorDicts.errorCode[UNKNOWN_ERROR] = &CommonError{
		ClientMessage:         "Unhandled error.",
		ServiceMessage:        "An unhandled error has occured. Please contact the developer.",
		OzzoValidationMessage: nil,
		ErrorCode:             UNKNOWN_ERROR,
	}

	errorDicts.httpCode[UNKNOWN_ERROR] = http.StatusInternalServerError

	return errorDicts
}
