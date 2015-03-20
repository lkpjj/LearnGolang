package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 3)
		ch1 <- "input"
	}()

	select {
	case t := <-ch1:
		fmt.Println(t)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
