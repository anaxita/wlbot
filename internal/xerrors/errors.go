package xerrors

import (
	"errors"
	"fmt"
)

var ErrValidate = errors.New("validation failed")
var ErrNotFound = errors.New("not found")
var ErrWrongInput = errors.New("wrong input")

// Wrap wraps text by err. If err is empty returns nil
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", text, err)
}
