package main

import "fmt"

type StructA struct {
	Name string
}
type StructB struct {
	Name string
}

func main() {
	a := StructA{}
	a.Print("i am a")
	fmt.Println(a)

	b := StructB{}
	b.Print("i am b")
	fmt.Println(b)

}

// func (receiver） funcName()
func (a *StructA) Print(content string) {
	fmt.Println("StructA:", content)
	a.Name = "AA"
}

func (a StructB) Print(content string) {
	fmt.Println("StructB:", content)
	a.Name = "BB"
}
