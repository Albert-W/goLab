package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//for i := 0; i < 1000; i++ {
	//	 //不加go 就无限循环0；
	//	 //func(i int) {
	//	 go func(i int) {
	//		for {
	//			fmt.Printf("Hello from "+
	//				"goroutine %d\n", i)
	//		}
	//	}(i)
	//}
	var a [10]int
	for i := 0; i < 10; i++ {
		 //不加go 就无限循环0；
		 //func(i int) {
		 // ii 的值 来自外面的i
		 go func(ii int) {
			for {
				a[ii]++
				runtime.Gosched()
			}
			//如果不定义i, index out of range
		}(i)
	}
	//没人交出控制权
	time.Sleep(time.Millisecond)
	//需要使用channal
	fmt.Println(a) //[1738 719 667 798 806 754 551 638 558 591]
}
