package main

import "fmt"

/*课堂作业
//
//如果匿名字段和外层结构有同名字段，应该如何进行操作？
//请思考并尝试。
*/

type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
	Sex  string
}
type student struct {
	human
	Name string
	Age  int
	Sex  string
}

func main() {
	homework_10()
}
func homework_10() {
	a := teacher{Name: "joe", Age: 60, human: human{0}}
	b := student{Name: "david", Age: 70, human: human{1}}
	fmt.Println(a, b)
	a.Name = "joe2"
	a.Age = 14
	a.human.Sex = 100
	b.Name = "david2"
	b.Age = 15
	b.Sex = "male"
	fmt.Println(a, b)

}
