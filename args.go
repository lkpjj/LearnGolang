package main

import "os"
import "fmt"

func main(){
	ags1 := os.Args
	ags2 := os.Args[1:]
	ags3 := os.Args[2]

	fmt.Println(ags1,ags2,ags3)
}
