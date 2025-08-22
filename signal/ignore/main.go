package main

import (
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signal.Ignore(syscall.SIGINT, syscall.SIGTERM)

	time.Sleep(10 * time.Second)
}
