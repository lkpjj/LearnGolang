package main

import "fmt"
import "os/exec"

func main(){
	lscmd := exec.Command("bash","-c", "ls -a -l -h")
	lsout,err := lscmd.Output()
	if err != nil {
panic(err)
}
	fmt.Println("> ll output is:")
fmt.Println(string(lsout))
}
