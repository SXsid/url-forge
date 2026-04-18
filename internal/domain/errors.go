package domain

import "errors"

var (
	ErrCodeNotFound = errors.New("short code not found")
	ErrURLExists    = errors.New("url already exists")
	ErrInternal     = errors.New("internal error")
)
