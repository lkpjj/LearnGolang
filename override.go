package main

import (
	"fmt"
)

//定义两种不同类型的结构体
type Person struct {
	Id   int
	Name string
}

type Student struct {
	Person
	Score int
}

func main() {
	p := Student{Person{Id: 1001, Name: "kevin"}, 100}
	p.test()
}

func (this Person) test() {
	fmt.Println("this is Person")
}

func (this Student) test() {
	fmt.Println("this is Student")
}
