package main

import "fmt"

/*
课堂作业

根据为结构增加方法的知识，尝试声明一个底层类型为int的类型，
并实现调用某个方法就递增100。

如：a:=0，调用a.Increase()之后，a从0变成100。
*/

type TZ2 int

func main() {
	var aTZ2 TZ2 = 0
	fmt.Println(aTZ2)
	aTZ2.Increase(100)
	fmt.Println(aTZ2)
	aTZ2.Increase(100)
	fmt.Println(aTZ2)

}

func (a *TZ2) Increase(num int) {
	*a += TZ2(num)
}
