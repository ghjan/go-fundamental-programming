package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type User2 struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User2
	title string
}

func (u User2) Status(name string, val int) string {
	return "Status " + name + ", my name is " + u.Name + ". I have " + strconv.Itoa(val) + " tokens"
}

func main() {
	fmt.Println("-----reflectionEmbededStructDemo----")
	reflectionEmbededStructDemo()
	fmt.Println("-----reflectionModifyDemo1----")
	reflectionModifyDemo1()
	fmt.Println("-----reflectionModifyDemo2----")
	reflectionModifyDemo2()
	fmt.Println("-----reflectionFuncinvokeDemo----")
	reflectionFuncinvokeDemo()
}

func reflectionEmbededStructDemo() {
	m := Manager{User2: User2{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	//取出匿名字段 通过索引
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}

func reflectionModifyDemo1() {
	x := 123
	v := reflect.ValueOf(&x) //注意这里是取地址
	fmt.Println(x)
	v.Elem().SetInt(999)
	fmt.Println(x)
}

func reflectionModifyDemo2() {
	u := User2{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	//if v.Kind() != reflect.Ptr || !v.Elem().CanSet()
	if (v.Kind() != reflect.Ptr && v.Kind() != reflect.Interface) || !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}

}

func reflectionFuncinvokeDemo() {
	u := User2{1, "OK", 12}
	fmt.Println(u.Status("joe", 200))

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Status")

	args := []reflect.Value{reflect.ValueOf("joe"), reflect.ValueOf(300)}

	fmt.Println(mv.Call(args)[0])

}
