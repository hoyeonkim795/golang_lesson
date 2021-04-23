package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var pattern string
	var diamondPattern string
	rightInput := false
	for !rightInput {
		fmt.Println("역직삼각형은 1, 다이아몬드는 2, 평행사변형은 3, 기울어진 삼각형은 4, 연속된 다이아몬드는 5")
		fmt.Print("원하는 패턴에 해당하는 번호를 입력하세요: ")
		fmt.Scanln(&pattern)

		switch pattern {
		case "1", "2", "3", "4" :
			rightInput = true
		case "5":
			for !rightInput{
				fmt.Println("두 번 연속 다이아몬드는 2, 세번 연속 다이아몬드는 3")
				fmt.Print("입력 조건에 맞게 원하는 번호를 입력하세요:")
				fmt.Scanln(&diamondPattern)
				err, convertedDiamondPattern:= checkRightInput(diamondPattern)
				if err != nil {
					continue
				}
				rightInput = true
				switch convertedDiamondPattern {
				case 2:
					pattern = "5"
				case 3:
					pattern = "6"
				default:
					rightInput = false
				}
			}
		default:
			fmt.Println("1,2,3,4,5 중에 입력하셔야 합니다.")
		}
	}
	rightInput = false
	var num string
	for !rightInput {
		fmt.Print("라인 수를 입력하세요 :")
		fmt.Scan(&num)
		err, convertedNum := checkRightInput(num)
		if err != nil {
			fmt.Println("숫자를 입력하세요")
			continue
		}
		rightInput = true
		switch pattern{
		case "1":
			triangle(convertedNum)
		case "2":
			if !diamondLineInput(convertedNum) {
				rightInput = false
				continue
			}
			diamond(convertedNum)
		case "3":
			parallelogram(convertedNum)
		case "4":
			obliqueTriangle(convertedNum)
		case "5":
			if !diamondLineInput(convertedNum) {
				rightInput = false
				continue
			}
			twiceDiamond(convertedNum)
		case "6":
			if !diamondLineInput(convertedNum) {
				rightInput = false
				continue
			}
			tripleDiamond(convertedNum)

		}
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


func diamond(convertedNum int) {
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

func twiceDiamond(convertedNum int) {
	for j := 1; j > -1; j-- {
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

func tripleDiamond(convertedNum int) {
	for j := 2; j > -1; j-- {
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