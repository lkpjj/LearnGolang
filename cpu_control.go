package main

import "time"
import "fmt"

func main() {

	// 首先我们看下基本的频率限制。假设我们得控制请求频率，
	// 我们使用一个通道来处理所有的这些请求，这里向requests
	// 发送5个数据，然后关闭requests通道
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(time.Millisecond * 200)

	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}

	busylimter := make(chan time.Time, 3)
	for i := 0; i <= 2; i++ {
		busylimter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Microsecond * 200) {
			busylimter <- t
		}
	}()

	busyrequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		busyrequest <- i
	}
	close(busyrequest)
	for req := range busyrequest {
		<-busyrequest
		fmt.Println("request", req, time.Now())
	}
}
