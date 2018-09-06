package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//range 也可以退出
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}

	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %d \n", id, n)
	//}
}
//chan<- 只能送数据 send;
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	//var c chan int // c == nil 定义了接受int 的channel;
	c := make(chan int)
	go func() {
		for{ //用死循环来收 channel
			n := <-c
			fmt.Println(n) // 1,2
		}

	}()
	c <- 1  //没有收会deadlock
	c <- 2
	time.Sleep(time.Millisecond)
	 //
	 //开了10个channel
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	//等待打印完毕
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	//缓冲区为3, 可以提升性能
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd' //dead lock
	time.Sleep(time.Millisecond)
}

// 永远是发送方来close
func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
