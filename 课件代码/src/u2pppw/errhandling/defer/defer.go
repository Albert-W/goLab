package main

import (
	"fmt"
	"os"

	"bufio"

	"u2pppw/functional/fib"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		//相当于有先进后出的栈
		//99, 98, ... , 0
		defer fmt.Println(i)
		if i == 30 {
			// Uncomment panic to see
			// how it works with defer
			// panic("printed too many")
		}
	}
}
// 1, 1,2 ,3 ,5,8,
func writeFile(filename string) {
	//0666是UNIX的权限；
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)
	//file, err :=os.Create(filename)
	//err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			//如果是pathError;
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	//最后关闭；就近添加，以防忘记
	defer file.Close()
	//带内存，比较比
	writer := bufio.NewWriter(file)
	//导入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
