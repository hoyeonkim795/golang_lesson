package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// sig라는 채널을 만든다
	sig := make(chan os.Signal, 1)
	// sig 채널에 지정한 signal을 보냄
	signal.Notify(sig, syscall.SIGINT)
	s := <-sig
	if s == syscall.SIGINT {
		fmt.Println("Thank you")

	}
}
