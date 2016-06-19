package unit

import (
	"errors"
	"fmt"
)

var ErrNotSet = errors.New("Field not specified")
var ErrNotExist = errors.New("Does not exist")
var ErrNotSupported = errors.New("Not supported yet")
var ErrUnknownType = errors.New("Unknown type")
var ErrPathNotAbs = errors.New("Path specified is not absolute")

type ParseError struct {
	Source string
	Err    error
}

func ParseErr(source string, err error) ParseError {
	return ParseError{
		Source: source,
		Err:    err,
	}
}

func (err ParseError) Error() string {
	return fmt.Sprintf("%s: %s", err.Source, err.Err)
}