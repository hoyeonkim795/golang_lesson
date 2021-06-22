package input

import (
	"myapp10/config"
	"fmt"
	"strconv"
)


func GetPatternRepeatCount(patternNum int) (int, bool){
	if patternNum == config.Diamond {
		repeatCount, isValid := GetValidInput("다이아몬드 패턴 카운트를 입력하세요",1,100)
		return repeatCount, isValid
	}
	return 1,true
}


func IsValidDiamondLineCount(lineCount int) bool{
	if lineCount %2 == 0{
		fmt.Println("다이아몬드 패턴을 선택했을 경우 라인의 수는 홀 수 여야 합니다.")
		return false
	}
	return true
}

func GetValidInput(question string, startRange int, endRange int) (int,bool){
	var input string
	for {
		fmt.Print("range(",startRange,"~",endRange, ") ", question)
		fmt.Scanln(&input)
		if getExitCommand(input) { // 종료 명령어 입력했을 경우
			return 0, false
		}
		convertedNum, err := convertStringToInt(input)
		if !isValidRange(convertedNum, err, startRange, endRange){
			continue
		}
		return convertedNum, true
	}
}


func convertStringToInt(input string) (int, error) {
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