package exit

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Interrupt(interruptChan chan os.Signal) {
	for {
		<-interruptChan
		fmt.Println("Thank you")
		os.Exit(0)
	}
}

func SignalHandle() chan os.Signal{
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}


