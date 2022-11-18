package xerrors

import (
	"errors"
	"fmt"
)

var (
	ErrValidate    = errors.New("validation failed")
	ErrNotFound    = errors.New("not found")
	ErrWrongInput  = errors.New("wrong input")
	ErrSendMessage = errors.New("send message")
	ErrHealthCheck = errors.New("health check")
)

// Wrap wraps text by err. If err is empty returns nil.
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", text, err)
}
