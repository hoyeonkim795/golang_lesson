package input

import (
	"strings"
)

func exitInput(input string) bool{
	switch strings.ToUpper(input) {
	case "Q", "QUIT":
		return true
	default:
		return false
	}
}

