package main

import "fmt"

// 익명함수
// 함수 선언 자체가 프로그래밍 전역으로 초기화되면서 메모리를 잡아먹는다.
// 기능을 수행할 때마다 함수를 찾아서 호출해야하기 때문이다.

func addDeclared(nums ...int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return
}

func main() {
	var nums = []int{10, 12, 13, 14, 16}

	addAnonymous := func(nums ...int) (result int) {
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		return
	}

	fmt.Println(addAnonymous(nums...))
	fmt.Println(addDeclared(nums...))
}

// 선언함수는 프로그램이 시작됨과 동시에 모두 읽는다.
// 하지만 익명함수는 해당 함수가 실행되는 곳에서 읽는다.
// 즉, 선언함수보다 익명함수가 나중에 읽힌다.
