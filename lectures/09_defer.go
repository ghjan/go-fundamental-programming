package main

import "fmt"

func main() {

	fmt.Println("======deferDemo1==")
	deferDemo1()

	fmt.Println("======deferDemo2==")
	deferDemo2()
	fmt.Println("======deferDemo3==")
	deferDemo3()
	fmt.Println("======deferDemo4==")
	deferDemo4()
}
func deferDemo1() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}

func deferDemo2() {

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

func deferDemo3() {

	for i := 0; i < 3; i++ {
		//defer一个闭包
		defer func() {
			fmt.Println(i)
		}()
	}
}

func deferDemo4() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

//Panic 可以在任何地方引发，但recover只有在defer调用的函数中有效
func B() {
	//defer 语句放在panic之前 先注册一个defer函数做recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err is: %s\n", err)
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")

}

func C() {
	fmt.Println("Func C")
}

