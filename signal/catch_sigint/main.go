package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	done := make(chan struct{})

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		<-sigChan
		done <- struct{}{}
	}()

	infiniteApp(done)

}

func infiniteApp(done <-chan struct{}) {
	cnt := 10
	for {
		select {
		case <-done:
			fmt.Println("stop!")
			return
		default:
			fmt.Println(cnt)
			time.Sleep(300 * time.Millisecond)
			cnt--
			if cnt == 0 {
				cnt = 10
			}
		}
	}
}
