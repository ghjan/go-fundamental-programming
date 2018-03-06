package main

import (
	"fmt"
)

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
	fmt.Println("===========嵌套结构nestDemo======")
	nestDemo()
	fmt.Println("===========匿名字段anonymouseSecDemo======")
	anonymouseSecDemo()
	fmt.Println("===========赋值assignDemo======")
	assignDemo()
	fmt.Println("===========嵌入当做继承inheritDemo======")
	inheritDemo()
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

//使用字面值定义并且初始化struct
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
	//匿名结构 定义结构 并且初始化
	a := &struct {
		Name string
		Age  int
	}{
		Name: "joeA",
		Age:  39,
	}
	fmt.Println(a)
}

//嵌套结构
func nestDemo() {
	type person2 struct {
		Name    string
		Age     int
		Contact struct {
			Phone, City string
		}
	}
	a := person2{}
	fmt.Println(a)

	b := person2{Name: "joe", Age: 49}
	b.Contact.Phone = "123456789"
	b.Contact.City = "Beijing"
	fmt.Println(b)
}

//匿名字段
func anonymouseSecDemo() {
	//初始化 如果不写字段名称 必须和结构定义时候的次序一致
	a := person{"joe", 19}
	fmt.Println(a)
}

//赋值
func assignDemo() {
	a := person{"joe", 59}
	var b person
	b = a

	fmt.Println(b)
	fmt.Println(b == a)

	c := person{"joe", 59}
	fmt.Println(a == c)

	//不同类型是不能比较的 编译错误
	//type person3 struct {
	//	Name string
	//	Age  int
	//}
	//d :=person3{"joe", 59}
	//fmt.Println(a==d)
}

//golang没有继承 使用嵌入结构模拟
func inheritDemo() {
	type human struct {
		Sex int
	}
	type teacher struct {
		human
		Name string
		Age  int
	}
	type student struct {
		human
		Name string
		Age  int
	}
	a := teacher{Name: "joe", Age: 60, human: human{0}}
	b := student{Name: "david", Age: 70, human: human{1}}
	fmt.Println(a, b)
	a.Name = "joe2"
	a.Age = 14
	a.human.Sex = 100
	b.Name = "david2"
	b.Age = 15
	b.Sex = 110
	fmt.Println(a, b)
}
