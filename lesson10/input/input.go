package input

import (
	"../patternList"
	"fmt"
	"strconv"
)
func PatternCount(pattern int) (int, bool){
	if pattern == 2 {
		patternCount, checkExit := GetInput("다이아몬드 패턴 카운트를 입력하세요",1,100)
		return patternCount, checkExit
	}
	return 1,false
}

func PrintPattern(pattern int, lineCount int, patternCount int) bool{
	switch pattern {
	case 1:
		fmt.Println(patternList.Triangle(lineCount, patternCount ))
	case 2:
		if !isValidDiamondLineCount(lineCount){
			return false
		}
		fmt.Println(patternList.Diamond(lineCount, patternCount))
	case 3:
		fmt.Println(patternList.Parallelogram(lineCount, patternCount))
	case 4:
		fmt.Println(patternList.ObliqueTriangle(lineCount, patternCount))
	}
	return true
}

func isValidDiamondLineCount(lineCount int) bool{
	if lineCount %2 == 0{
		fmt.Println("다이아몬드 패턴을 선택했을 경우 라인의 수는 홀 수 여야 합니다.")
		return false
	}
	return true
}

func GetInput(question string, startRange int, endRange int) (int,bool){
	var input string
	for {
		fmt.Print("range(",startRange,"~",endRange, ") ", question)
		fmt.Scanln(&input)
		if exitInput(input) {
			return 0, true
		}
		convertedNum, err := convertStringToInt(input)
		if !isValidRange(convertedNum, err, startRange, endRange){
			continue
		}
		return convertedNum, false
	}
}

func convertStringToInt(input string) (int, error) {
	// 입력 받은 값에 대해서 실행
	convertedNum, err := strconv.Atoi(input)
	return convertedNum, err
}

func isValidRange(convertedNum int, err error, startRange int, endRange int) bool{
	if err != nil || convertedNum < startRange || convertedNum > endRange{
		fmt.Println("입력 조건에 맞는 값을 입력해 주세요.")
		return false
	}
	return true
}