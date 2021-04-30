package main

import (
	"./exit"
	"./input"
	"fmt"
)


func main() {
	var (
		checkExit bool
		pattern int
		lineCount int
		patternCount int
	)
	//signal handle
	sig := exit.SignalHandle()
	go exit.Interrupt(sig)

	for {
		// 패턴 입력받기
		fmt.Println("역직삼각형은 1, 다이아몬드는 2, 평행사변형은 3, 기울어진 삼각형은 4")
		pattern, checkExit = input.AskInput("원하는 패턴에 해당하는 번호를 입력하세요:",1,4)
		if checkExit{
			break
		}

		patternCount, checkExit = input.PatternCount(pattern)

		if checkExit{
			break
		}

		// 라인 입력받기
		for {
			lineCount, checkExit = input.AskInput("라인의 수를 입력하세요 :",1,100)
			if checkExit {
				break
			}
			if input.PrintPattern(pattern, lineCount, patternCount) {
				break
			}
		}
		if checkExit {
			break
		}

	}
}



