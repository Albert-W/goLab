package main

import (
	"fmt"
	//"runtime"
	"math/cmplx"
)
func variable(){
	var a int =0
	var s string ="abc"
	fmt.Print(a,s)
}

func enums()  {
	
}

func main() {
	//go func() {
	//	fmt.Println("1")
	//}()
	//go func() {
	//	fmt.Println("2")
	//}()
	//go func() {
	//	fmt.Println("3")
	//}()
	//runtime.Gosched()
	variable()
	c := 3+4i
	fmt.Println(cmplx.Abs(c))

	

}
