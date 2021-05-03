package main

import (
	"./config"
	"./exit"
	"./input"
	"./print"
	"fmt"
)


func main() {
	var (
		isValid bool
		patternNum int
		lineCount int
		repeatCount int
	)

	sig := exit.SignalHandle()
	go exit.Interrupt(sig)

	for {
		fmt.Println("역직삼각형은 1, 다이아몬드는 2, 평행사변형은 3, 기울어진 삼각형은 4")
		patternNum, isValid = input.GetValidInput("원하는 패턴에 해당하는 번호를 입력하세요:",config.PatternNumStartRange,config.PatternNumEndRange)
		if !isValid{
			return
		}
		repeatCount, isValid = input.GetPatternRepeatCount(patternNum)
		if !isValid{
			return
		}

		// if else ?
		switch patternNum {
		case config.Diamond:
			for {
				lineCount, isValid = input.GetValidInput("라인의 수를 입력하세요 :",config.InputStartRange,config.InputEndRange)
				if input.IsValidDiamondLineCount(lineCount) {
					break
				}
			}
		default:
			lineCount, isValid = input.GetValidInput("라인의 수를 입력하세요 :",config.InputStartRange,config.InputEndRange)
		}

		if !isValid{
			return
		}

		print.Pattern(patternNum, lineCount, repeatCount)
	}
}



