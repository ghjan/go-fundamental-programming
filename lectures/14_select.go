package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----chanSelectDemo---select用于接收chan消息-")
	chanSelectDemo()
	fmt.Println("-----chanSelectDemo2 select用于发送chan消息----")
	chanSelectDemo2()
}
func chanSelectDemo() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("c1 is closed")
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2 is closed")
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"
	close(c1)
	//close(c2)
	for i := 0; i < 2; i++ {
		<-o
	}
}

func chanSelectDemo2() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case c <- 0:
			fmt.Println("i is ", i)
		case c <- 1:
			fmt.Println("i is ", i)
		}
	}
}
