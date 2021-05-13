package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)


func main() {
	//signal handle
	sig := signalHandle()
	go interrupt(sig)
	for {
		// 패턴 입력받기
		pattern, checkExit := inputPattern()
		if checkExit{
			break
		}
		// 라인 입력받기
		checkExit = inputPrintLine(pattern)
		if checkExit {
			break
		}
	}
}

func interrupt(interruptChan chan os.Signal) {
	for {
		<-interruptChan
		fmt.Println("Thank you")
		os.Exit(0)
	}
}

func signalHandle() chan os.Signal{
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}

func inputPattern() (string, bool){
	var pattern string
	for {
		fmt.Println("역직삼각형은 1, 다이아몬드는 2, 평행사변형은 3, 기울어진 삼각형은 4")
		fmt.Print("원하는 패턴에 해당하는 번호를 입력하세요: ")
		fmt.Scanln(&pattern)

		switch pattern {
		case "1", "2", "3", "4" :
			return pattern, false
		default:
			if exitInput(pattern){
				return pattern, true
			}
			fmt.Println("1,2,3,4 중에 입력하셔야 합니다.")
		}
	}
}


func triangle(lineNum int) {
	for i := lineNum; i > 0; i-- {
		stars := strings.Repeat("*", i)
		fmt.Println(stars)
	}
}
func diamond(lineNum int, patternCount int) {
	for j := patternCount - 1; j > -1; j-- {
		for i := 1; i < lineNum+1; i = i+2 {
			blank := strings.Repeat(" ", (lineNum-i)/2)
			stars := strings.Repeat("*", i)
			fmt.Println(blank,stars,blank)
		}
		for i := lineNum-2; i > 1; i = i-2 {
			blank := strings.Repeat(" ", (lineNum-i)/2)
			stars := strings.Repeat("*", i)
			fmt.Println(blank,stars,blank)
		}
	}
	blank := strings.Repeat(" ", (lineNum-1)/2)
	stars := strings.Repeat("*",1)
	fmt.Println(blank,stars,blank)
}

func parallelogram(lineNum int)  {
	for i :=lineNum -1; i > -1; i-- {
		blank := strings.Repeat(" ", i)
		stars := strings.Repeat("*", lineNum)
		fmt.Println(blank, stars)
	}
}

func obliqueTriangle(lineNum int) {
	for i :=lineNum ; i > 0; i-- {
		blank := strings.Repeat(" ", i-1)
		stars := strings.Repeat("*", i)
		fmt.Println(blank, stars)
	}
}

func inputPrintLine(pattern string) bool{
	rightInput := false
	for !rightInput {
		lineNum, checkExit:= askInput("라인의 수를 입력하세요 :")
		if checkExit{
			return true
		}
		rightInput = true
		switch pattern {
		case "1":
			triangle(lineNum)
		case "2":
			if lineNum % 2 == 0{
				fmt.Println("다이아몬드 패턴을 선택했을 경우 라인의 수는 홀 수 여야 합니다.")
				rightInput = false
				continue
			}
			patternCount,checkExit := askInput("pattern count 를 입력하세요")
			if checkExit{
				return true
			}
			diamond(lineNum, patternCount)
		case "3":
			parallelogram(lineNum)
		case "4":
			obliqueTriangle(lineNum)
		}
	}
	return false
}

func askInput(question string) (int,bool){
	var input string
	for {
		fmt.Print("range(1~100) ", question)
		fmt.Scanln(&input)
		if exitInput(input) {
			return 0,true
		}
		convertedNum, err := strconv.Atoi(input)
		if err != nil || convertedNum < 1 || convertedNum > 100{
			fmt.Println("입력 조건에 맞는 값을 입력해 주세요.")
			continue
		}
		return convertedNum, false
	}
}

func exitInput(input string) bool{
	switch strings.ToUpper(input) {
	case "Q", "QUIT":
		return true
	default:
		return false
	}
}



