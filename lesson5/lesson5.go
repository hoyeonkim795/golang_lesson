package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var pattern string

	for {
		fmt.Print("원하는 패턴에 해당하는 번호를 입력하세요. (역직삼각형은 1, 다이아몬드는 2) : ")
		fmt.Scanln(&pattern)
		if pattern == "1" || pattern == "2"{
			break
		} else {
			fmt.Println("1과 2 중에 입력하셔야 합니다.")
		}
	}

	var num string
	for {
		fmt.Print("라인 수를 입력하세요 :")
		fmt.Scan(&num)

		if convertedNum, err:= strconv.Atoi(num); err == nil{
			if convertedNum < 1 {
				fmt.Println("0보다 큰 숫자를 입력하세요")
				continue
			}
			if pattern == "2" && convertedNum % 2 == 0 {
				fmt.Println("다이아몬드 패턴을 선택했을 경우 라인의 수는 홀수 여야 합니다.")
				continue
			}
			// 역직삼각형 패턴
			if pattern == "1"{
				triangle(convertedNum)
				break
			} else if pattern == "2" { // 다이어몬드 패턴
				diamond(convertedNum)
				break
			}
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
	for i := 1; i < convertedNum+1; i++ {
		blank := strings.Repeat(" ", (convertedNum-i)/2)
		stars := strings.Repeat("*", i)
		fmt.Println(blank,stars,blank)
		i += 1
	}
	for i := convertedNum-2; i > 0; i-- {
		blank := strings.Repeat(" ", (convertedNum-i)/2)
		stars := strings.Repeat("*", i)
		fmt.Println(blank,stars,blank)
		i -= 1
	}
}