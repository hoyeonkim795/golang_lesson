package print

import (
	"../config"
	"../patternList"
	"fmt"
)

func Pattern(patternNum int, lineCount int, repeatCount int) bool{
	switch patternNum {
	case config.ReversedTriangle:
		fmt.Println(patternList.ReversedTriangle(lineCount))
	case config.Diamond:
		fmt.Println(patternList.Diamond(lineCount, repeatCount))
	case config.Parallelogram:
		fmt.Println(patternList.Parallelogram(lineCount))
	case config.ObliqueTriangle:
		fmt.Println(patternList.ObliqueTriangle(lineCount))
	}
	return true
}