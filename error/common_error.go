package errors

import (
	"encoding/json"
	"fmt"
	"github.com/a666hn/go-libs/converter"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type OzzoValidationError map[string]string

type CommonError struct {
	ClientMessage         string              `json:"message"`
	ServiceMessage        interface{}         `json:"data"`
	OzzoValidationMessage OzzoValidationError `json:"validation,omitempty"`
	ErrorCode             ErrorCode           `json:"code"`
	ErrorMessage          *string             `json:"-"`
	ErrorTrace            *string             `json:"-"`
}

func (ce *CommonError) Error() string {
	return fmt.Sprintf(
		"CommonError: %v. Trace: %v",
		converter.PointerToString(ce.ErrorMessage),
		converter.PointerToString(ce.ErrorTrace),
	)
}

func (ce *CommonError) SetClientMessage(msg string) {
	ce.ClientMessage = msg
}

func (ce *CommonError) SetServiceMessage(msg interface{}) {
	ce.ServiceMessage = msg
}

func (ce *CommonError) SetValidationDtoMessage(message interface{}) {
	if ee, ok := message.(validation.Errors); ok {
		ce.OzzoValidationMessage = ce.buildOzzoValidationMessage(ee)
	}
}

func (ce *CommonError) buildOzzoValidationMessage(dtoError validation.Errors) OzzoValidationError {
	var message OzzoValidationError
	errorMessage, _ := dtoError.MarshalJSON()
	_ = json.Unmarshal(errorMessage, &message)

	return message
}

func NewCommonError(
	errCode ErrorCode,
	err error,
) *CommonError {
	var (
		message        *string
		trace          *string
		clientMessage              = "Unhandled error."
		serviceMessage interface{} = "An unhandled error has occured. Please contact the developer."
		commonError                = errorDictionaries.errorCodes[errCode]
	)

	if err != nil {
		s := err.Error()
		message = converter.StringToPointer(s)

		ss := fmt.Sprintf("%+v", err)
		trace = converter.StringToPointer(ss)
	}

	if commonError == nil {
		return &CommonError{
			ClientMessage:         clientMessage,
			ServiceMessage:        serviceMessage,
			OzzoValidationMessage: nil,
			ErrorCode:             errCode,
			ErrorMessage:          message,
			ErrorTrace:            trace,
		}
	}

	if _err, ok := err.(*CommonError); ok {
		return _err
	}

	return &CommonError{
		ClientMessage:         commonError.ClientMessage,
		ServiceMessage:        commonError.ServiceMessage,
		OzzoValidationMessage: commonError.OzzoValidationMessage,
		ErrorCode:             errCode,
		ErrorMessage:          message,
		ErrorTrace:            trace,
	}
}
