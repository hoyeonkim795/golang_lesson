package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var pattern string
	rightInput := false
	for !rightInput {
		fmt.Println("역직삼각형은 1, 다이아몬드는 2, 평행사변형은 3, 기울어진 삼각형은 4")
		fmt.Print("원하는 패턴에 해당하는 번호를 입력하세요: ")
		fmt.Scanln(&pattern)

		switch pattern {
		case "1", "2", "3", "4" :
			rightInput = true
		default:
			fmt.Println("1,2,3,4,5 중에 입력하셔야 합니다.")
		}
	}

	rightInput = false
	for !rightInput {
		convertedNum := askInput("라인의 수를 입력하세요 :")
		rightInput = true
		switch pattern {
		case "1":
			triangle(convertedNum)
		case "2":
			if convertedNum %2 == 0 {
				fmt.Println("다이아몬드는 홀수 라인만 입력할 수 있습니다.")
				rightInput = false
				continue
			}
			patternCount := askInput("pattern count 를 입력하세요")
			diamond(convertedNum, patternCount)
		case "3":
			parallelogram(convertedNum)
		case "4":
			obliqueTriangle(convertedNum)
		}
	}
}

func askInput(question string) int{
	var input string
	for {
		fmt.Print(question)
		fmt.Scanln(&input)
		err, convertedNum := checkRightInput(input)
		if err != nil {
			fmt.Println("입력 조건에 맞는 값을 입력해 주세요.")
			continue
		}
		return convertedNum
	}
}

func checkRightInput(num string) (error, int){
	convertedNum, err := strconv.Atoi(num)
	return err, convertedNum
}


func triangle(convertedNum int) {
	for i := convertedNum; i > 0; i-- {
		stars := strings.Repeat("*", i)
		fmt.Println(stars)
	}
}

func diamondLineInput(convertedNum int) bool {
	if convertedNum %2 == 0 {
		fmt.Println("다이아몬드 패턴을 선택했을 경우 라인의 수는 홀수 여야 합니다.")
		return false
	}
	return true
}


func diamond(convertedNum int, patternCount int) {
	for j := patternCount - 1; j > -1; j-- {
		for i := 1; i < convertedNum+1; i = i+2 {
			blank := strings.Repeat(" ", (convertedNum-i)/2)
			stars := strings.Repeat("*", i)
			fmt.Println(blank,stars,blank)
		}
		for i := convertedNum-2; i > j; i = i-2 {
			blank := strings.Repeat(" ", (convertedNum-i)/2)
			stars := strings.Repeat("*", i)
			fmt.Println(blank,stars,blank)
		}
	}
}

func parallelogram(convertedNum int)  {
	for i :=convertedNum -1; i > -1; i-- {
		blank := strings.Repeat(" ", i)
		stars := strings.Repeat("*", convertedNum)
		fmt.Println(blank, stars)
	}
}

func obliqueTriangle(convertedNum int) {
	for i :=convertedNum ; i > 0; i-- {
		blank := strings.Repeat(" ", i-1)
		stars := strings.Repeat("*", i)
		fmt.Println(blank, stars)
	}
}