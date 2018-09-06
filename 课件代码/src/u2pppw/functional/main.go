package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"u2pppw/functional/fib"
)
//函数类型；
type intGen func() int
// g 为接收者；给函数实现Reader接口
func (g intGen) Read(
	p []byte) (n int, err error) {
		//取得下一个元素；
	next := g()
	//结束条件
	if next > 10000 {
		return 0, io.EOF
	}
	// 把next写成字符，换行；
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	//当文件来使用
	printFileContents(f)
}
