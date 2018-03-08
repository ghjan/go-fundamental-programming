package main

import (
	"fmt"
	"strconv"
)

type (
	byteAlias int8
	rune int32
	文本   string
)

func main() {
	var b 文本
	b = "中文啊亲"
	fmt.Println("文本b：" + b)

	var ii uint8 = 255
	fmt.Printf("ii:%d\n", ii)

	//	.\03_type-alias.go:18:16: constant 255 overflows byteAlias
	var bb byteAlias = 127
	fmt.Printf("bb:%d\n", bb)

	var aaa float32 = 1.1
	ba := int(aaa)
	fmt.Printf("ba:%d\n", ba)
	var aa2 int = 2
	ba2 := float32(aa2)
	fmt.Printf("ba2:%f\n", ba2)

	var aa3 int = 65
	ba3 := string(aa3)
	fmt.Println(ba3)

	ba4 := strconv.Itoa(aa3)
	fmt.Println(ba4)

}
