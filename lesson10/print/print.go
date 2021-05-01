package print

import (
	"../patternList"
	"fmt"
)

func PrintPattern(patternNum int, lineCount int, repeatCount int) bool{
	switch patternNum {
	case 1:
		fmt.Println(patternList.Triangle(lineCount, repeatCount ))
	case 2:
		fmt.Println(patternList.Diamond(lineCount, repeatCount))
	case 3:
		fmt.Println(patternList.Parallelogram(lineCount, repeatCount))
	case 4:
		fmt.Println(patternList.ObliqueTriangle(lineCount, repeatCount))
	}
	return true
}