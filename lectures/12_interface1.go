package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

func main() {
	fmt.Println("------interfaceDemo1----")
	interfaceDemo1()
	fmt.Println("------interfaceDemo2----")
	interfaceDemo2()

}

func interfaceDemo1() {
	var a USB
	a = PhoneConnecter{"name of PhoneConnecter"}
	//a.name = "" // compile error
	a.Connect()
}

func interfaceDemo2() {
	//var a USB
	a := PhoneConnecter{"name of PhoneConnecter"}
	a.Connect()
	a.name += " in interfaceDemo2" // ok
	a.Connect()
	//验证一下a是否USB
	Disconnect(a)
}

func Disconnect(usb USB){
	fmt.Println("Disconnected.")
}
