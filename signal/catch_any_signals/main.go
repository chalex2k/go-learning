package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)

	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGINT:
			fmt.Println("SIGINT")
		case syscall.SIGTERM:
			fmt.Println("SIGTERM")
		default:
			fmt.Println(sig)
		}
	}
}
