package main

import "fmt"

type StructA3 struct {
	//私有成员对于同一个 package里面还是可见的
	name string
}

func main() {
	//struct字段的访问权限和package的关系
	var a StructA3
	a.Print()
	fmt.Println(a.name)
}

func (a *StructA3) Print() {
	a.name = "123"
	fmt.Println(a.name)
}
