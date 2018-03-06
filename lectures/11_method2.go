package main

import "fmt"

type TZ int

func main() {
	//method value
	var a TZ
	a.Print("i am method value")

	//method expression
	//直接通过类型而不是变量调用方法
	(*TZ).Print(&a, "i am method expression")
}

//类型别名也可以自定义方法
//这也是为什么类型别名也需要强制转换的原因
func (a *TZ) Print(content string) {
	fmt.Println("TZ:", content)
}
