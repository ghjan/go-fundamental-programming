package main

import "fmt"

func main() {
	a, b := functionDemo1(1, 2)
	fmt.Println(a, b)
	fmt.Println("=====functionDemo2 按值传参===")
	//按值传参
	functionDemo2(a, b)
	fmt.Println(a, b)
	fmt.Println("=====functionDemo2 reference传参===")
	s1 := []int{1, 2, 3, 4}
	functionDemo3(s1)
	fmt.Println(s1)

	//函数也可以作为一种类型使用
	aa := functionDemo1
	aa(3, 4)

	fmt.Println("=====匿名函數===")
	//匿名函數
	a3 := func() {
		fmt.Println("anonymouse function")
	}
	a3()

	fmt.Println("======闭包==")
	//闭包
	f := closure1(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

}

func functionDemo1(a int, b int) (int, int) {

	return a, b
}

//按值传参 int类型不定长度参数
func functionDemo2(s ...int) {
	s[0] = 5
	s[1] = 6
	fmt.Println(s)
}

//reference传参  slice
func functionDemo3(s []int) {
	s[0] = 5
	s[1] = 6
	fmt.Println(s)
}

func closure1(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
