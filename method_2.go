package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{age: 2, name: "one"}
	p.test(1)

	//将方法直接赋给另一个函数 f1
	var f1 func(int) = p.test
	f1(2)

	person.test(p, 3)
}

//定义一种方法
func (this person) test(x int) {
	fmt.Println("age:", this.age, "name:", this.name)
	fmt.Println("x=", x)
}
