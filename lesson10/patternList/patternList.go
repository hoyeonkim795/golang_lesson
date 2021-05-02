package patternList

import (
	"../config"
	"strings"
)

func Triangle(lineCount int) string{
	var resultPattern string
	for lineCountIdx := lineCount; lineCountIdx > 0; lineCountIdx-- {
		stars := strings.Repeat(config.Star, lineCountIdx)
		resultPattern += stars + "\n"
	}
	return resultPattern
}

func Diamond(lineCount int, repeatCount int) string{
	var resultPattern string
	for repeatCountIdx := repeatCount - 1; repeatCountIdx > -1; repeatCountIdx-- {
		for top := 1; top < lineCount+1; top = top+2 {
			blanks := strings.Repeat(config.Blank, (lineCount-top)/2)
			stars := strings.Repeat(config.Star, top)
			resultPattern += blanks + stars + blanks +"\n"
		}
		for bottom := lineCount-2; bottom > 1; bottom = bottom-2 {
			blanks := strings.Repeat(config.Blank, (lineCount-bottom)/2)
			stars := strings.Repeat(config.Star, bottom)
			resultPattern += blanks + stars + blanks + "\n"
		}
	}
	blanks := strings.Repeat(config.Blank, (lineCount-1)/2)
	stars := strings.Repeat(config.Star,1)
	resultPattern += blanks + stars + "\n"
	return resultPattern
}

func Parallelogram(lineCount int)  string{
	var resultPattern string
	for lineCountIdx :=lineCount -1; lineCountIdx > -1; lineCountIdx-- {
		blanks := strings.Repeat(config.Blank, lineCountIdx)
		stars := strings.Repeat(config.Star, lineCount)
		resultPattern += blanks + stars + "\n"
	}
	return resultPattern
}

func ObliqueTriangle(lineCount int) string{
	var resultPattern string
	for lineCountIdx :=lineCount ; lineCountIdx > 0; lineCountIdx-- {
		blanks := strings.Repeat(config.Blank, lineCountIdx-1)
		stars := strings.Repeat(config.Star, lineCountIdx)
		resultPattern += blanks + stars + "\n"
	}
	return resultPattern
}
