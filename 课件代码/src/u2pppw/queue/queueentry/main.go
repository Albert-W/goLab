package main

import (
	"fmt"
	"u2pppw/queue"
)

func main() {
	q := queue.Queue{1} //原始值

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
