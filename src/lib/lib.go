package lib

import (
	"errors"
)

func CheckRequiredFields(fields []string) (err error) {
	for _, field := range fields {
		if field == "" {
			err = errors.New("this field cannot be empty")
		}
	}
	if err != nil {
		return err
	}
	return nil
}
