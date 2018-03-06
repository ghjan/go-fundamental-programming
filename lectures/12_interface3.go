package main

import "fmt"

type Painter interface {
	Draw()
}
type CowBoyer interface {
	DrawTheGun()
}

type XiaoWang struct {
	name string
}

func (self *XiaoWang) Draw() {
	fmt.Println("i can draw")
}
func (self *XiaoWang) DrawTheGun() {
	fmt.Println("i can draw the gun")
}

func main() {
	var xm interface{} = new(XiaoWang)

	pt := xm.(Painter)
	pt.Draw()

	cb := xm.(CowBoyer)
	cb.DrawTheGun()
}
