package main

import (
	"fmt"
)

func main() {
	chs := make(chan string)
	go func() { chs <- "message" }()
	value := <-chs
	fmt.Println(value)
}
