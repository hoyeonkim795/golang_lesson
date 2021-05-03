package patternList

import (
	"../config"
	"bytes"
	"strings"
)

func ReversedTriangle(lineCount int) string{
	var resultPattern bytes.Buffer
	for lineCountIdx := lineCount; lineCountIdx > 0; lineCountIdx-- {
		stars := strings.Repeat(config.Star, lineCountIdx)
		resultPattern.WriteString(stars)
		resultPattern.WriteString("\n")
	}
	return resultPattern.String()
}

func Diamond(lineCount int, repeatCount int) string{
	var resultPattern bytes.Buffer
	for repeatCountIdx := repeatCount - 1; repeatCountIdx > -1; repeatCountIdx-- {
		for top := 1; top < lineCount+1; top = top+2 {
			blanks := strings.Repeat(config.Blank, (lineCount-top)/2)
			stars := strings.Repeat(config.Star, top)
			resultPattern.WriteString(blanks)
			resultPattern.WriteString(stars)
			resultPattern.WriteString(blanks)
			resultPattern.WriteString("\n")
		}
		for bottom := lineCount-2; bottom > 1; bottom = bottom-2 {
			blanks := strings.Repeat(config.Blank, (lineCount-bottom)/2)
			stars := strings.Repeat(config.Star, bottom)
			resultPattern.WriteString(blanks)
			resultPattern.WriteString(stars)
			resultPattern.WriteString(blanks)
			resultPattern.WriteString("\n")

		}
	}
	blanks := strings.Repeat(config.Blank, (lineCount-1)/2)
	stars := strings.Repeat(config.Star,1)
	resultPattern.WriteString(blanks)
	resultPattern.WriteString(stars)
	resultPattern.WriteString("\n")
	return resultPattern.String()
}

func Parallelogram(lineCount int)  string{
	var resultPattern bytes.Buffer
	for lineCountIdx :=lineCount -1; lineCountIdx > -1; lineCountIdx-- {
		blanks := strings.Repeat(config.Blank, lineCountIdx)
		stars := strings.Repeat(config.Star, lineCount)
		resultPattern.WriteString(blanks)
		resultPattern.WriteString(stars)
		resultPattern.WriteString("\n")
	}
	return resultPattern.String()
}

func ObliqueTriangle(lineCount int) string{
	var resultPattern bytes.Buffer
	for lineCountIdx :=lineCount ; lineCountIdx > 0; lineCountIdx-- {
		blanks := strings.Repeat(config.Blank, lineCountIdx-1)
		stars := strings.Repeat(config.Star, lineCountIdx)
		resultPattern.WriteString(blanks)
		resultPattern.WriteString(stars)
		resultPattern.WriteString("\n")
	}
	return resultPattern.String()
}
