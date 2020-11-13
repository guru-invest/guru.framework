package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

func Throw(err error, message string) error{
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}()
	return errors.Wrap(err, message)
}