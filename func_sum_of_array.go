package main

import "fmt"

// 这个函数的返回值为两个int
func vals(a [5]int) (int, int) {
    sum := 0
    for _,n := range a{
        sum += n;
    }
    return len(a),sum
}

func main() {

    s := [5]int{1,2,3,4,5}
    n,sum := vals(s)
    fmt.Println(n,sum)
}