package main

import (
	"fmt"
)

//递归实现加法运算
func sum(a int) int {
	if 1 == a {
		return 1
	}

	return a + sum(a-1)
}

func main() {
	fmt.Println(sum(5))
}
