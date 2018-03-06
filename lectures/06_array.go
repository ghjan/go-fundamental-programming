package main

import "fmt"

func main() {
	arrayDemo()

}
func arrayDemo() {
	var a [2]int
	var b [2]int
	b = a
	fmt.Println(b)
	a2 := [3]int{1, 2}
	fmt.Println(a2)
	a3 := [...]int{19: 1}
	fmt.Println(a3)
	a4 := [...]int{0: 1, 1: 2, 2: 3}
	fmt.Println(a4)
	a5 := [...]int{4, 5, 6}
	fmt.Println(a5)
	//数组指针
	var p *[20]int = &a3
	fmt.Println(p)
	fmt.Println(*p)
	//指针数组
	x, y := 1, 2
	ap := [...]*int{&x, &y}
	fmt.Println(ap)

	//new关键字
	pn := new([10]int)
	fmt.Println(pn)

	//多维数组
	//非顶级的维度必须指定长度
	am := [...][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(am)
}
