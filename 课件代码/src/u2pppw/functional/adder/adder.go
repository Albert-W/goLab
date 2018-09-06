package main

import "fmt"

// 无参，返加为匿名函数；
func adder() func(int) int {
	//自由变量
	sum := 0
	//v 局部变量，函数可以使用sum。
	return func(v int) int {
		sum += v
		return sum
	}
}
//递归的定义
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	a := adder() //is trivial and also works.
	for i := 0;i<10;i++{
		//fmt.Printf("0+1+...+%d = %d\n", i, a(i))
		//println(a(i))
	}
	//println(a(10)) //10
	//println(a(7))  //17
	fmt.Printf("%T\n", a(10))
	//println(a())	//not enough arguments in call to a have () want (int)
	//println(a(13)) //30
	//println(adder())
	//println(adder())
	//println(adder())
	//println(a)
	//println(a)
	//println(adder())
	//println(adder())
	//println(&a)
	//a := adder2(0)
	////打印0加到9的部分和
	//for i := 0; i < 10; i++ {
	//	//s 是部分和；
	//	var s int
	//	s, a = a(i)
	//	fmt.Printf("0 + 1 + ... + %d = %d\n",
	//		i, s)
	//}
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
	fmt.Println(&f) // "16"
	fmt.Println(&f) // "16"
	fmt.Println(&f) // "16"
}
