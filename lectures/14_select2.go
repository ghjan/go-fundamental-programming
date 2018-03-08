package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-----chanSelectDemo3 select 可设置超时----")
	chanSelectDemo3()
	fmt.Println("-----chanSelectDemo4 select 可设置超时----")
	chanSelectDemo4()
	fmt.Println("-----tickerDemo 计时器----")
	tickerDemo()
}

func chanSelectDemo3() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("after 3 seconds")
		break
	}
}

func chanSelectDemo4() {
	c := make(chan bool)
	go func() {
		c <- true
	}()
LOOP:
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(3 * time.Second):
			fmt.Println("chanSelectDemo4 after 3 seconds")
			break LOOP
		}
	}
}

//time.Ticker 结构体，这个对象以指定的时间间隔重复的向通道 C 发送时间值
func tickerDemo() {
	tick := time.Tick(1e8)
	boom := time.After(5e8+1)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
