package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var num string
	for {
		fmt.Print("숫자를 입력하세요 :")
		fmt.Scan(&num)
		if convertedNum, err := strconv.Atoi(num); err == nil {
			if convertedNum > 0 && convertedNum < 101 {
				for i := 1; i < convertedNum+1; i++ {
					stars := strings.Repeat("*", i)
					fmt.Println(stars)
				}
				break
			}
		}
	}

}