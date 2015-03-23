package main
import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
)

func check(e error){
	if nil != e{
		panic(e)
	}

	fmt.Println("this is no error")
}

func main(){
	str := []byte("你好，世界！")
	err := ioutil.WriteFile("/tmp/dat", str,0644)
	check(err)
	
	f,err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

// 你可以写入字节切片
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)
 
    // 也可以使用`WriteString`直接写入字符串
    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)
 
    // 调用Sync方法来将缓冲区数据写入磁盘
    f.Sync()

 // `bufio`除了提供上面的缓冲读取数据外，还
    // 提供了缓冲写入数据的方法
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)
 
    // 使用Flush方法确保所有缓冲区的数据写入底层writer
    w.Flush()
}
