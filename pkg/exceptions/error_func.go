package exceptions

import "fmt"

func ErrRequired(msg string) error {
	return fmt.Errorf("%s is required", msg)
}

func ErrNotFound(msg string) error {
	return fmt.Errorf("%s not found", msg)
}
