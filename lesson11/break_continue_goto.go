package main

import "fmt"

func main() {
	i := 0

TEST1:
	for {
		if i == 0 {
			break TEST1
		}
	}
	fmt.Println("END")
	// END 출력

	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d", i)
	}
	//13579111315

	var num int
	fmt.Print("자연수를 입력하세요")
	fmt.Scanln(&num)

	if num == 1 {
		goto One
	} else if num == 2 {
		goto Two
	} else {
		goto OTHER
	}

One:
	fmt.Printf("1을 입력했군")
	goto END
Two:
	fmt.Printf("2를 입력했군")
	goto END
OTHER:
	fmt.Printf("다른걸 입력했군")
END:
}
