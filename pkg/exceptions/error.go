package exceptions

import "fmt"

var (
	ErrInternalServer = fmt.Errorf("internal server error")
	ErrUnauthorized   = fmt.Errorf("unauthorized")
	ErrForbidden      = fmt.Errorf("forbidden")
)
