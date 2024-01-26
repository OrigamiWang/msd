package errx

import (
	"fmt"
)

type ErrX interface {
	Code() int
	Error() string
	String() string
}

type errX struct {
	code int
	msg  string
}

func (e *errX) Code() int {
	return e.code
}

func (e *errX) Error() string {
	return e.msg
}

func (e errX) String() string {
	return fmt.Sprintf("&errX{code:%d;msg:\"%s\"}", e.code, e.msg)
}

func New(code int, msg string) ErrX {
	return &errX{code, msg}
}

func IsErrX(e interface{}) bool {
	_, ok := e.(ErrX)
	return ok
}

func IsError(e interface{}) bool {
	_, ok := e.(error)
	return ok
}
