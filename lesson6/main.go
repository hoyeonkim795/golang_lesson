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
			fmt.Println("1,2,3,4 중에 입력하셔야 합니다.")
		}
	}

	rightInput = false
	var num string
	for !rightInput {
		fmt.Print("라인 수를 입력하세요 :")
		fmt.Scan(&num)
		convertedNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("숫자를 입력하세요")
			continue
		}
		rightInput = true
		switch pattern{
		case "1":
			triangle(convertedNum)
		case "2":
			if convertedNum%2 == 1 {
				diamond(convertedNum)
				continue
			}
			rightInput = false
			fmt.Println("다이아몬드 패턴을 선택했을 경우 라인의 수는 홀수 여야 합니다.")
		case "3":
			parallelogram(convertedNum)
		case "4":
			obliqueTriangle(convertedNum)
		}
	}
}

func triangle(convertedNum int) {
	fmt.Println("역 직삼각형 패턴입니다.")
	for i := convertedNum; i > 0; i-- {
		stars := strings.Repeat("*", i)
		fmt.Println(stars)
	}
}

func diamond(convertedNum int)  {
	fmt.Println("다이아몬드 패턴입니다.")
	for i := 1; i < convertedNum+1; i = i+2 {
		blank := strings.Repeat(" ", (convertedNum-i)/2)
		stars := strings.Repeat("*", i)
		fmt.Println(blank,stars,blank)
	}
	for i := convertedNum-2; i > 0; i = i-2 {
		blank := strings.Repeat(" ", (convertedNum-i)/2)
		stars := strings.Repeat("*", i)
		fmt.Println(blank,stars,blank)
	}
}

func parallelogram(convertedNum int)  {
	fmt.Println("평행사변형 패턴입니다.")
	for i :=convertedNum -1; i > -1; i-- {
		blank := strings.Repeat(" ", i)
		stars := strings.Repeat("*", convertedNum)
		fmt.Println(blank, stars)
	}
}

func obliqueTriangle(convertedNum int) {
	fmt.Println("기울어진 삼각형 패턴입니다.")
	for i :=convertedNum ; i > 0; i-- {
		blank := strings.Repeat(" ", i-1)
		stars := strings.Repeat("*", i)
		fmt.Println(blank, stars)
	}
}
