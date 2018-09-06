package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com
email2 99@qq.com
email3 wu@163.com
`

func main()  {
	//exp := `[a-zA-Z0-9]+@.+\..+`
	//exp := `[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`
	exp := `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`
	compile:= regexp.MustCompile(exp)
	//match := compile.FindString(text)
	//match := compile.FindAllString(text,-1)
	match := compile.FindAllStringSubmatch(text,-1)
	for _, m :=range match{
		fmt.Println(m)
	}
	//fmt.Println(match)
}
