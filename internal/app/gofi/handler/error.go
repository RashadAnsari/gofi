package handler

import "errors"

var (
	ErrInvalidHeaderSyntax = errors.New("invalid header syntax")
	ErrInvalidQuerySyntax  = errors.New("invalid query syntax")
)
