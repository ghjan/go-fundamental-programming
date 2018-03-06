package main

import (
	"fmt"
)

//我们可以说任何类型都实现了empty接口
type empty interface {
}

type USB2 interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}
type PhoneConnecter2 struct {
	name string
}

//只实现了Connecter，没有实现USB2
type TVConnecter struct {
	name string
}

func (tv TVConnecter) Connect() {
	fmt.Println("Connected:", tv.name)
}

func (pc PhoneConnecter2) Name() string {
	return pc.name
}
func (pc PhoneConnecter2) Connect() {
	fmt.Println("Connected:", pc.name)
}

func main() {
	fmt.Println("------interfaceDemo1_2----")
	interfaceDemo1_2()
	fmt.Println("------interfaceDemo2_2----")
	interfaceDemo2_2()
	fmt.Println("------interfaceTypeSwitchDemo----")
	interfaceTypeSwitchDemo()
	fmt.Println("------interfaceConvertDemo----")
	interfaceConvertDemo()
	//fmt.Println("------interfaceConvertErrorDemo----")
	//interfaceConvertErrorDemo()
	fmt.Println("------interfaceConvertDemo2 ----")
	interfaceConvertDemo2()
	fmt.Println("------nilDemo----")
	nilDemo()
}

func interfaceDemo1_2() {
	var a USB2
	a = PhoneConnecter2{"name of PhoneConnecter2"}
	//a.name = "" // compile error
	a.Connect()
}

func interfaceDemo2_2() {
	//var a USB2
	a := PhoneConnecter2{"name of PhoneConnecter2"}
	a.Connect()
	a.name += " in interfaceDemo2_2" // ok
	a.Connect()
	//验证一下a是否USB2
	Disconnect_2(a)
}

func interfaceTypeSwitchDemo() {
	//var a USB2
	a := PhoneConnecter2{"name of PhoneConnecter2"}
	a.Connect()
	a.name += " in interfaceTypeSwitchDemo" // ok
	a.Connect()
	//验证一下a是否USB2
	Disconnect3(a)
}

func Disconnect_2(usb USB2) {
	//类型判断 ok pattern
	if pc, ok := usb.(PhoneConnecter2); ok {
		fmt.Println("Disconnected:", pc.name)
	} else {
		fmt.Println("Unknown Disconnected.")
	}
}

//空接口 type switch
func Disconnect3(usb interface{}) {
	//类型判断
	switch v := usb.(type) {
	case PhoneConnecter2:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device disconnected.")
	}
}

//接口相互转换  只能够超集-》子集  USB->Connecter
func interfaceConvertDemo() {
	pc := PhoneConnecter2{"PhoneConnecter2"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	//a.Name() //编译错误 a不是USB 仅仅是Connecter
}

//cannot convert tv (type TVConnecter) to type USB2:
//TVConnecter does not implement USB2 (missing Name method)
//func interfaceConvertErrorDemo() {
//	tv := TVConnecter{"TVConnecter"}
//	var a USB2
//	a = USB2(tv)
//	a.Connect()
//}
//将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
func interfaceConvertDemo2() {
	pc := PhoneConnecter2{"PhoneConnecter2"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	pc.name = "PC"
	//a是pc的复制品，修改了pc，对a没有影响
	a.Connect()
}

//只有当接口存储的类型和对象都为nil时，接口才等于nil
func nilDemo() {
	var a interface{}
	fmt.Println(a == nil)

	var p *int = nil
	a = p
	fmt.Println(a == nil)
}
