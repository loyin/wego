package validations

import (
	"errors"
	"strings"
)

func Valid(obj interface{}) error {
	valid := Validation{}
	b, _ := valid.Valid(obj)
	if !b {
		var msg string
		for _, err := range valid.Errors {
			m := strings.Join([]string{err.Field, err.Message}, ":")
			msg = strings.Join([]string{m, msg}, ";")
		}
		return errors.New(msg)
	}
	return nil
}
