package main

import (
	"./exit"
	"./input"
	"fmt"
)


func main() {
	var (
		isValid bool
		patternNum int
		lineCount int
		repeatCount int
	)
	//signal handle
	sig := exit.SignalHandle()
	go exit.Interrupt(sig)

	for {
		// 패턴 입력받기
		fmt.Println("역직삼각형은 1, 다이아몬드는 2, 평행사변형은 3, 기울어진 삼각형은 4")
		patternNum, isValid = input.GetValidInput("원하는 패턴에 해당하는 번호를 입력하세요:",1,4)
		// 입력한 patternNum값이 종료 명령어 인지 확인
		if !isValid{
			break
		}
		// 패턴 반복 횟수 (repeatCount) 입력 받기
		repeatCount, isValid = input.GetPatternRepeatCount(patternNum)
		// 입력한 patternNumcount가 종료 명령어 인지 확인
		if !isValid{
			break
		}

		switch patternNum {
		case 2:
			for {
				lineCount, isValid = input.GetValidInput("라인의 수를 입력하세요 :",1,100)
				if input.IsValidDiamondLineCount(lineCount) {
					break
				}
			}
		default:
			lineCount, isValid = input.GetValidInput("라인의 수를 입력하세요 :",1,100)
		}

		if !isValid{
			break
		}

		input.PrintPattern(patternNum, lineCount, repeatCount)
	}
}



