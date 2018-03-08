package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	/*
		 多个goroutime保持同步,两种方法:
		使用有缓存的chan
		使用sync.WaitGroup
	*/
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("----chanDemo5-使用有缓存的chan--")
	chanDemo5()
	fmt.Println("----chanDemo6--使用sync.WaitGroup-")
	chanDemo6()
}

//使用有缓存的chan
func chanDemo5() {
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go GO(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func GO(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	//time.Sleep(2 * time.Second)
	c <- true
}

//使用sync.WaitGroup
func chanDemo6() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go GO2(&wg, i)
	}
	wg.Wait()
}

func GO2(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}
