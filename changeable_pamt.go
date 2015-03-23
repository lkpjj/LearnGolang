package main

import (
	"fmt"
)

//可变参数函举例
func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 5))

	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))
}
