package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	//加锁
	lock sync.Mutex
}

//+1；
func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() {
		//加锁
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
