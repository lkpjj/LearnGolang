package main

import (
	"fmt"
	"strconv"
)

func main() {
	buffer := make(chan string, 10)
	flag := make(chan bool)

	go func() {
		for {
			msg, ok := <-buffer
			if ok {
				fmt.Println("There is elements: ", msg)
			} else {
				fmt.Println("buffer is empty")
				flag <- true
				return
			}
		}
	}()

	for i := 0; i < 4; i++ {
		buffer <- "msg-->" + strconv.Itoa(i)
		fmt.Println("sending jobs")
	}

	close(buffer)
	fmt.Println("send all jobs")
	<-flag
}
