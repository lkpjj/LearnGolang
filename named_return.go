package main

import (
	"fmt"
)

func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y, 1
}

func main() {
	fmt.Println(split(10))
}
