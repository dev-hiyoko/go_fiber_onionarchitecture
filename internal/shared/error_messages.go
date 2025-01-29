package shared

import "errors"

type Error string

const (
	NoneCode = "000"

	DefaultError              Error = "bad request"
	ValidateError             Error = "validate error"
	UserEmailExistsError      Error = "email already exists"
	UserOriginalIDExistsError Error = "original ID already exists"
	UserPasswordNotMatchError Error = "password does not match"
)

var ErrorCodes = map[string]string{
	DefaultError.String():              NoneCode,
	ValidateError.String():             "001",
	UserEmailExistsError.String():      "002",
	UserOriginalIDExistsError.String(): "003",
	UserPasswordNotMatchError.String(): "004",
}

func (e Error) Error() error {
	return errors.New(e.String())
}

func (e Error) String() string {
	return string(e)
}

func GetErrorCodes(err error) string {
	if val, ok := ErrorCodes[err.Error()]; ok {
		return val
	}

	return NoneCode
}
