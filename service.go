package main

import (
	"context"
	"errors"
)

// StringService ...
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct{}

// ErrEmpty is returned when input sstring is empty
var ErrEmpty = errors.New("Empty string")
