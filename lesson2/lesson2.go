package main

import (
	"fmt"
	"strings"
)

func main() {
	for i:=1; i<6; i++ {
		stars := strings.Repeat("*",i)
		fmt.Println(stars)
	}
}