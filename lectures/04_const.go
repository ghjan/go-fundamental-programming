package main

import "fmt"

const a int = 1
const b = 'A'
const (
	text   = "123"
	length = len(text)
	num    = b * 20
)

const s = "abc"
const (
	aa = len(s)
	bb
	cc
)

const (
	a3, b3 = 1, "2"
	c3, d3
)

const (
	a4 = 'A'
	b4 = iota
	c4 = 'a'
	d4 = iota
)
const (
	MONDAY = iota
	TUESDAY
)

const (
	_  = iota
	KB float64= 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func operatorWei() {
	//位运算符
	/*
					二进制   十进制
					01100   12
					10110   22
		-----------------------
		二元位运算符  二进制   十进制    逻辑
			&       00100   4        相同位的两个数字­都为1，则为1；若有一个不为1，则为0。
			|       11110   30       相同位只要一个为1即为1。否则为0
			^       11010   26       相同位不同为1则为1，否则为0。
			&^      01000   8        如果第二个数是1，则强制把第一个数改成0，否则不变。
	*/
	fmt.Println(12 & 22)
	fmt.Println(12 | 22)
	fmt.Println(12 ^ 22)
	fmt.Println(12 &^ 22)
}

func homework_4() {
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}
func main() {
	fmt.Println(aa, bb, cc)
	fmt.Println(a3, b3, c3, d3)
	fmt.Println(a4, b4, c4, d4)
	fmt.Println(MONDAY, TUESDAY)

	//取反
	fmt.Println(^8)
	//位移运算符
	fmt.Println(1 << 10 << 10)
	operatorWei()

	ai := 3
	if ai > 0 && (10/a) > 1 {
		fmt.Println(10 / a)
	} else {
		fmt.Println(ai)
	}

	homework_4()
}
