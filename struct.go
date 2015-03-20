package main

import (
	"fmt"
)

//定义一个新的结构体类型
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"kevin", 25})

	fmt.Println(person{name: "tom", age: 10})

	//未初始化的整型变量默认为0，string类型的变量默认为空
	fmt.Println(person{age: 10})

	fmt.Println(&person{"bob", 10})

	s1 := person{"cat", 10}
	fmt.Println(s1.name, s1.age)

	s2 := &s1
	fmt.Println(s2.age)

	s2.age = 100
	fmt.Println(s1.age)
}
