package main

import "fmt"

func main() {
	fmt.Println("======homework_9==")
	homework_9()

}

func homework_9() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i) //值拷贝
		defer func() {
			fmt.Println("defer_closure i=", i)
		}()
		fs[i] = func() { fmt.Println("closure i=", i) }
	}

	for _, f := range fs {
		f()
	}
}
