package main

import (
	"./exit"
	"./input"
)


func main() {
	//signal handle
	sig := exit.SignalHandle()
	go exit.Interrupt(sig)
	for {
		// 패턴 입력받기
		pattern, checkExit := input.Pattern()
		if checkExit{
			break
		}
		// 라인 입력받기
		checkExit = input.PrintPattern(pattern)
		// checkExit을 반복해서 안쓰고 하는 방법?
		if checkExit {
			break
		}
	}
}



