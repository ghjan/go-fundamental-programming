package main

import (
	"net"
	"time"
)

func main() {
	serversWorkerDemo()
}
func serversWorkerDemo() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	ch := make(chan string)
	go pool3(ch, 4)
	go server3(l, ch)
	time.Sleep(10 * time.Second)
}

func server3(l net.Listener, ch chan string) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler3(c, ch)
	}
}

func handler3(c net.Conn, ch chan string) {
	addr := c.RemoteAddr().String()
	ch <- addr
	time.Sleep(100 * time.Millisecond)
	c.Write([]byte("ok"))
	c.Close()
}

func pool3(ch chan string, n int) {
	wch := make(chan int)
	results := make(chan int)
	for i := 0; i < n; i++ {
		go logger3(wch, results)
	}
	go parse(results)
	for {
		addr := <-ch
		l := len(addr)
		wch <- l
	}
}

func logger3(wch chan int, results chan int) {
	for {
		data := <-wch
		data++
		results <- data
	}
}

func parse(results chan int) {
	for {
		<-results
	}
}
