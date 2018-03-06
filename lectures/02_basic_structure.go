// 当前程序的包名
package main

// 导入其它的包 别名
import std "fmt"

// 常量的定义
const PI = 3.14
const (
	const1 = "1"
	const2 = 2
	const3 = 3
)

// 全局变量的声明与赋值
var name = "gopher"

var (
	name1 = "1"
	name2 = 2
	name3 = 3
)

// 一般类型声明
type newType int

type (
	type1 float32
	type2 string
	type3 byte
)

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由 main 函数作为程序入口点启动
func main() {
	std.Println("Hello world!你好，世界！")
	var bbb int = 1
	std.Printf("bbb:%d\n", bbb)
	cc := 2
	std.Printf("c:%d\n", cc)
	var a, b, c, d = 1, 2, 3, 4
	std.Printf("%d,%d,%d,%d\n", a, b, c, d)

	//方法体内可以使用:=, 可以代替var关键字
	ia, ib := 5, 6
	std.Printf("%d,%d\n", ia, ib)



}
