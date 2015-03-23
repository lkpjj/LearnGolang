package main

import "fmt"
import "os"
import "strings"

func main(){
	os.Setenv("Foo", "hehe")
	fmt.Println(os.Getenv("Foo"))

	fmt.Println()
	for _,e := range os.Environ(){
		pair := strings.Split(e, "+")
fmt.Println(pair[0])
}
}
