package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int
	for {
		fmt.Print("숫자를 입력하세요 :")
		fmt.Scanln(&num)
		if 0 < num  && num < 101 {
			for i := 1; i < num+1; i++ {
				stars := strings.Repeat("*", i)
				fmt.Println(stars)
			}
			break
		}
	}
}