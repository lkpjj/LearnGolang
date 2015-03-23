package main

import (
	"fmt"
)

func main() {
	chs := make(chan string, 5)
	chs <- "one"
	chs <- "two"
	fmt.Println(<-chs)
	fmt.Println(<-chs)
}
