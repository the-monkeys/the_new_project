package common

import "errors"

var (
	NotFound   = errors.New("not found")
	BadRequest = errors.New("bad request")
	Internal   = errors.New("internal server error")
)

const (
	DATE_TIME_FORMAT = "2006-01-02T15:04:05Z07:00"
)
