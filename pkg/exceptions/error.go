package exceptions

import "fmt"

var (
	ErrInternalServer = fmt.Errorf("internal server error")
)
