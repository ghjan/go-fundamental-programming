package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//fmt.Println("----helloCSP---------")
	//helloCSP()
	//fmt.Println("----timerDemo---------")
	//timerDemo()
	fmt.Println("----pingpongDemo---------")
	pingpongDemo(2)
	fmt.Println("----pingpong3Demo---------")
	pingpongDemo(3)
	fmt.Println("----pingpong100Demo---------")
	pingpongDemo(100)

}
func timerDemo() {
	for i := 0; i < 8; i++ {
		c := timer(1 * time.Second)
		<-c
		fmt.Println("timer:", i)
	}
}
func helloCSP() {
	// create new channel of type int
	ch := make(chan int)
	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
		fmt.Println("hello CSP")
	}()
	// read from channel
	<-ch
}

func timer(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func pingpongDemo(count int) {
	var Ball int
	table := make(chan int)
	up := int(math.Max(2, float64(count)))
	for i := 0; i < up; i++ {
		go player(table)
	}

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int) {
	for {
		ball := <-table
		fmt.Println("ball:", ball)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
