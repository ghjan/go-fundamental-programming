package main

import "fmt"

func main() {
	closureGood()
	closureTrap()
}

//闭包的坑
func closureTrap() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	select {}
}


func closureGood() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	select {}
}
