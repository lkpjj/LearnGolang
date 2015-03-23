package main
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	dat,err := ioutil.ReadFile("/tmp/dat")
	check(err)
    fmt.Print(string(dat))

	f ,err := os.Open("/tmp/dat")
	b1 := make([]byte,5)
	n1,err := f.Read(b1)
    check(err)
	fmt.Printf("%d bytes:%s\n", n1,string(b1))
// 你也可以使用`Seek`来跳转到文件中的一个已知位置，并从
    // 那个位置开始读取数据
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
 
    // `io`包提供了一些帮助文件读取的函数。例如上面的方法如果
    // 使用方法`ReadAtLeast`函数来实现，将使得程序更健壮
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
 
    // 没有内置的rewind方法，但是可以使用`Seek(0,0)`来实现
    _, err = f.Seek(0, 0)
    check(err)
 
    // `bufio`包提供了缓冲读取文件的方法，这将使得文件读取更加
    // 高效
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))
 
    // 最后关闭打开的文件。一般来讲这个方法会在打开文件的时候，
    // 使用defer来延迟关闭
    f.Close()
}


