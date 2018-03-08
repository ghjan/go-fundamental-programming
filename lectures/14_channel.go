package main

import "fmt"

func main() {
	fmt.Println("----chanDemo---")
	chanDemo()
	fmt.Println("----chanDemo2---")
	chanDemo2()
	fmt.Println("----chanDemo3---")
	chanDemo3()
	fmt.Println("----chanDemo4---")
	chanDemo4()

}
func chanDemo() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!!!")
		c <- true   //向c写入值
	}()
	//堵塞了 等待能够从c中读取值
	<-c

}

func chanDemo2() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!!!")
		c <- true
		close(c) //必须的 否则会发生 deadlock死锁
	}()
	for v := range c {
		fmt.Println(v)
	}
}

//无缓冲的channel，要注意“取”先于“放”
func chanDemo3() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!!!")
		<-c
	}()
	c <- true
}

//有缓存异步 有缓冲的channel，要注意“放”先于“取”
func chanDemo4() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("GO GO GO!!!")
		c <- true
	}()
	<-c
}

