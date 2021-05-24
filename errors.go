package main

import (
	"errors"
	"fmt"
)

type ErrGet struct {
	StatusCode int
	Status     string
}

func (e *ErrGet) Error() string {
	return fmt.Sprintf("pricewatch: server failure has occurred: %d %s", e.StatusCode, e.Status)
}

var ErrUnsupportedHost = errors.New("pricewatch: unsupported host")
