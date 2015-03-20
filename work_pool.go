package main

import (
	"fmt"
	"time"
)

//从jobs接收工作，处理后把结果交给result
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 5

	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	for m := 1; m <= 9; m++ {
		jobs <- m
	}
	close(jobs)

	// 然后我们从results里面获得结果
	for a := 1; a <= 9; a++ {
		fmt.Println(<-results)
	}
}
