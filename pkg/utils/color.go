package utils

import (
	"fmt"
	"strings"
)

const (
	Black   = "black"
	Red     = "red"
	Green   = "green"
	Yellow  = "yellow"
	Blue    = "blue"
	Magenta = "magenta"
	Cyan    = "cyan"
	White   = "white"
)

func Color(color string, message string) string {
	const (
		black   = "\033[30m"
		red     = "\033[31m"
		green   = "\033[32m"
		yellow  = "\033[33m"
		blue    = "\033[34m"
		magenta = "\033[35m"
		cyan    = "\033[36m"
		white   = "\033[37m"
	)

	var (
		result string
		reset  = "\033[0m"
	)
	switch strings.ToLower(color) {
	case "black":
		result = black
	case "red":
		result = red
	case "green":
		result = green
	case "yellow":
		result = yellow
	case "blue":
		result = blue
	case "magenta":
		result = magenta
	case "cyan":
		result = cyan
	case "white":
		result = white
	default:
		result = black
	}
	result = fmt.Sprintf("%s%s%s", result, message, reset)
	return result
}
