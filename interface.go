package main

import (
	"fmt"
	"math"
)

//定义基本的几何方法,只声明而不具体实现
type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (c circle) area() float64 {
	return c.radius * 2 * math.Pi
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	// 这里circle和square都实现了geometry接口，所以
	// circle类型变量和square类型变量都可以作为measure
	// 函数的参数
	measure(s)
	measure(c)
}
