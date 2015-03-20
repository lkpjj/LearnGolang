package main

import (
	"fmt"
)

//首先类似于模板，定义函数类型为Callback
type Callback func(x, y float64) float64

func main() {
	x, y := 1.0, 20.0
	fmt.Println("add: ", test(x, y, add))
	fmt.Println("dvid ", test(x, y, dvid))
}

//提供接口，让别的函数去实现
func test(x, y float64, callback Callback) float64 {
	return callback(x, y)
}

func add(x, y float64) float64 {
	return x + y
}

func dvid(x, y float64) float64 {
	return x / y
}
