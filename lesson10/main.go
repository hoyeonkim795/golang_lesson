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
		pattern, checkExit = input.GetInput("원하는 패턴에 해당하는 번호를 입력하세요:",1,4)
		// 입력한 pattern값이 종료 명령어 인지 확인
		if checkExit{
			break
		}
		// patterncount 입력 받기
		patternCount, checkExit = input.PatternCount(pattern)
		// 입력한 patterncount가 종료 명령어 인지 확인
		if checkExit{
			break
		}
		// linecount 입력받기
		for {
			lineCount, checkExit = input.GetInput("라인의 수를 입력하세요 :",1,100)
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



