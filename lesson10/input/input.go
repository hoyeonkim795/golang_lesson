package input

import (
	"../patternList"
	"fmt"
	"strconv"

)

func Pattern() (int, bool){
	var pattern int
	var checkExit bool
	for {
		fmt.Println("역직삼각형은 1, 다이아몬드는 2, 평행사변형은 3, 기울어진 삼각형은 4")
		pattern, checkExit = askInput("원하는 패턴에 해당하는 번호를 입력하세요:",1,4)
		if checkExit{
			return 0, true
		}
		return pattern, false
	}
}


func PrintPattern(pattern int) bool{
	var (
		rightInput bool
		lineCount int
		patternCount int
		checkExit bool
	)

	rightInput = false
	for !rightInput {
		lineCount, checkExit = askInput("라인의 수를 입력하세요 :",1,100)
		if checkExit{
			return true
		}
		rightInput = true
		switch pattern {
		case 1:
			patternList.Triangle(lineCount)
		case 2:
			// 수정하고 싶은 부분 line을 입력 받는 부분을 따로 분리할 수 가 없음
		if lineCount % 2 == 0{
				fmt.Println("다이아몬드 패턴을 선택했을 경우 라인의 수는 홀 수 여야 합니다.")
				rightInput = false
				continue
			}
			patternCount,checkExit = askInput("pattern count 를 입력하세요",1,100)
			if checkExit{
				return true
			}
			patternList.Diamond(lineCount, patternCount)
		case 3:
			patternList.Parallelogram(lineCount)
		case 4:
			patternList.ObliqueTriangle(lineCount)
		}
	}
	return false
}

func askInput(question string, startRange int, endRange int) (int,bool){
	var input string
	for {
		fmt.Print("range(",startRange,"~",endRange, ") ", question)
		fmt.Scanln(&input)
		if exitInput(input) {
			return 0, true
		}
		convertedNum, err := strconv.Atoi(input)
		if err != nil || convertedNum < startRange || convertedNum > endRange{
			fmt.Println("입력 조건에 맞는 값을 입력해 주세요.")
			continue
		}
		return convertedNum, false
	}
}