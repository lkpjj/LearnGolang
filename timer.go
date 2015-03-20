package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 30)

	go func() {
		<-timer1.C
		fmt.Println("timer 1 expired")
	}()

	stop := timer1.Stop()
	if stop {
		fmt.Println("timer1 have cancled")
	}

}
