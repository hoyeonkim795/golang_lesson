package patternList

import (
	"fmt"
	"strings"
)

func Triangle(lineCount int) {
	for i := lineCount; i > 0; i-- {
		stars := strings.Repeat("*", i)
		fmt.Println(stars)
	}
}
func Diamond(lineCount int, patternCount int) {
	for j := patternCount - 1; j > -1; j-- {
		for i := 1; i < lineCount+1; i = i+2 {
			blank := strings.Repeat(" ", (lineCount-i)/2)
			stars := strings.Repeat("*", i)
			fmt.Println(blank,stars,blank)
		}
		for i := lineCount-2; i > 1; i = i-2 {
			blank := strings.Repeat(" ", (lineCount-i)/2)
			stars := strings.Repeat("*", i)
			fmt.Println(blank,stars,blank)
		}
	}
	blank := strings.Repeat(" ", (lineCount-1)/2)
	stars := strings.Repeat("*",1)
	fmt.Println(blank,stars,blank)
}

func Parallelogram(lineCount int)  {
	for i :=lineCount -1; i > -1; i-- {
		blank := strings.Repeat(" ", i)
		stars := strings.Repeat("*", lineCount)
		fmt.Println(blank, stars)
	}
}

func ObliqueTriangle(lineCount int) {
	for i :=lineCount ; i > 0; i-- {
		blank := strings.Repeat(" ", i-1)
		stars := strings.Repeat("*", i)
		fmt.Println(blank, stars)
	}
}
