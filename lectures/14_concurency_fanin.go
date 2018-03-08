package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("----faninDemo----------")
	faninDemo()
}

func faninDemo() {
	ch  := make(chan string )
	out := make(chan string)
	go producer(ch, "producer1", 100*time.Millisecond, 20)
	go producer(ch, "producer2", 250*time.Millisecond, 10)
	go reader(out)

	transfer(ch, out)
	//for i := range ch {
	//	out <- i
	//}
}

func transfer(in chan string, out chan string) {
	for {
		x, ok := <-in
		if !ok {
			return
		}
		out <- x
	}
}

func producer(ch chan string, title string, d time.Duration, limit int) {
	var i int
	for {
		if i > limit {
			return
		}
		ch <- title + ":" + strconv.Itoa(i)
		i++
		time.Sleep(d)
	}
}

func reader(out chan string) {
	for {
		x, ok := <-out
		if !ok {
			return
		}
		fmt.Println(x)
	}
}
