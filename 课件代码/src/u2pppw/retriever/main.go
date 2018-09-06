package main

import (
	"fmt"

	"time"

	"u2pppw/retriever/mock"
	"u2pppw/retriever/real"
)
//接口
type Retriever interface {
	Get(url string) string
}
//第二个接口
type Poster interface {
	Post(url string,
		form map[string]string) string
}
//需要接上http://
const url = "http://www.imooc.com"
//函数，参数r;
func download(r Retriever) string {
	return r.Get(url)
}
//接口对应的函数；
func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}
//定义一个组合接口
type RetrieverPoster interface {
	Retriever
	Poster
	//也可以添加其他方法；
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	//结构体变量；
	mockRetriever := mock.Retriever{
		Contents: "this is a fake imooc.com"}
	r = &mockRetriever
	//fmt.Println(download(r)) //

	inspect(r) //mock.Retriever {this is a fake imooc.com}

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r) //real.Retriever &{Mozilla/5.0 1m0s}

	// Type assertion， 也是判断方法
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	//查看类型和value;
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	//查看类型
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
