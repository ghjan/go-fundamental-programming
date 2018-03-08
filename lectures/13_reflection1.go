package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Status world")
}
func main() {
	var x float64 = 3.4
	Info(x)
	u := User{1, "OK", 12}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	fmt.Println(reflect.ValueOf(o))

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("error:not struct type")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fileds:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", f.Name, f.Type, val)
	}
	fmt.Println("Methods:")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}
