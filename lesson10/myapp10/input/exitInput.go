package input

import (
	"strings"
)

func getExitCommand(input string) bool{
	switch strings.ToUpper(input) {
	case "Q", "QUIT":
		return true
	default:
		return false
	}
}

