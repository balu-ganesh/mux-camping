package server

import "errors"

var ErrEmptyRequestBody = errors.New("request body is empty")
var ErrNoSlotID = errors.New("no slot id parameter in the URL")
var ErrNoUserID = errors.New("no user id parameter in the URL")

type BadRequestError struct {
	err error
}

func (e *BadRequestError) Error() string {
	return e.err.Error()
}

func (e *BadRequestError) Unwrap() error { return e.err }
