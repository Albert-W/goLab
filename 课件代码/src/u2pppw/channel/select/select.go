package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//随机sleep一断时间
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int //slice of int
	//计时器， channel time;
	tm := time.After(10 * time.Second)
	//定时，第秒看一个长度
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int  //nil
		var activeValue int
		if len(values) > 0 {  //有值
			activeWorker = worker
			activeValue = values[0]
		}
		//select 进行调度（任务分发）
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		//每次select 之间的时间，超时
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		//每秒看quene的长度
		case <-tick:
			fmt.Println(
				"queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
