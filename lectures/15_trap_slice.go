package main

import "fmt"

func main() {
	s := make([]int, 0) //初始容量为0 加入一个元素以后就会重新分配地址
	fmt.Println(s)
	pingpongTrap(s)
	fmt.Println(s)
	s = pingpongGood(s)
	fmt.Println(s)
}

//错误的版本 slice的陷阱
func pingpongTrap(s []int) {
	s = append(s, 3)
}

//正确的版本 希望修改slice 建议返回他
func pingpongGood(s []int) []int {
	s = append(s, 3)
	return s
}
