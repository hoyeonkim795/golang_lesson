package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	check := false
	for !check{
		fmt.Print("종료 명령어 입력 하세요. :")
		fmt.Scan(&input)
		convertedInput := convertUpper(input)
		check = checkConvertedInput(convertedInput)
	}
}

func convertUpper(input string) string{
	convertedInput := strings.ToUpper(input)
	return convertedInput
}

func checkConvertedInput(convertedInput string) bool {
	switch convertedInput {
	case "Q":
		return true
	case "QUIT":
		return true
	default:
		return false
	}
}