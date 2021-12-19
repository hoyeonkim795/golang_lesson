package main

import "fmt"

// 클로저는 함수 안에서 익명 함수를 정의해서 바깥쪽 함수에 선언한 변수에도 접근할 수 있는 함수

func next() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := next()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := next()
	fmt.Println(newInt())
}
