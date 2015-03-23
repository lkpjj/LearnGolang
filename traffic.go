package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	// signal := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		msg <- "you have a new msg!"
	}()

	select {
	case str := <-msg:
		fmt.Println("new msg", str)
	case <-time.After(time.Second * 3):
		fmt.Println("three mins have been used.")
		// default:
		// fmt.Println("no msg reveived")
	}
}
