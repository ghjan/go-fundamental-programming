package main

import "fmt"

type test struct {
}

type person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("===========structDemo1======")
	structDemo1()
	fmt.Println("===========structDemo2======")
	structDemo2()
	fmt.Println("===========structDemo3======")
	structDemo3()
	fmt.Println("===========匿名函数structAnonymouseDemo======")
	structAnonymouseDemo()
}

func structDemo1() {
	a := test{}
	fmt.Println(a)
	a2 := person{}
	fmt.Println(a2)
}

func structDemo2() {
	a := person{
		Name: "magic",
		Age:  19,
	}
	fmt.Println(a)
	a.Name = "joe"
	a.Age = 13
	fmt.Println(a)

	changePerson1(a)
	fmt.Println(a)    //changePerson1不能修改a
	changePerson2(&a) //传递地址
	fmt.Println(a)    //changePerson2修改了a
}

func changePerson1(per person) { //值拷贝
	per.Age = 23
	fmt.Println("changePerson1, per:", per)
}

func changePerson2(per *person) { //使用指针传递参数
	per.Age = 25
	fmt.Println("changePerson2, per:", per)
}

func changePerson3(per *person) { //使用指针传递参数
	per.Age = 28
	fmt.Println("changePerson3, per:", per)
}

func structDemo3() {
	a := &person{ //推荐使用地址
		Name: "magic",
		Age:  19,
	}
	fmt.Println(a)
	a.Name = "joe"
	a.Age = 13
	fmt.Println(a)

	changePerson2(a)
	fmt.Println(a)
	changePerson3(a)
	fmt.Println(a)
}

func structAnonymouseDemo() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "joeA",
		Age:  39,
	}
	fmt.Println(a)
}
