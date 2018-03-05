package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=------mapDemo----")
	mapDemo()
	fmt.Println("=------mapDemo2----")
	mapDemo2()
	fmt.Println("=------iterationDemo1----")
	iterationDemo1()
	fmt.Println("=------iterationDemo2----")
	iterationDemo2()
	fmt.Println("=------sortDemo----")
	sortDemo()
	fmt.Println("=------homework_8----")
	homework_8()

}
func mapDemo() {
	// map[keytype]valuetype
	m := make(map[int]string)
	m[1] = "OK"
	a := m[1]
	fmt.Println(m)
	fmt.Println(a)
	delete(m, 1)
	fmt.Println("After delete key from map")
	fmt.Println(m)
	fmt.Println(a)

}

func mapDemo2() {
	// map[keytype]valuetype
	m := make(map[int]map[int]string)
	m[1] = make(map[int]string)
	m[1][1] = "OK"
	a, ok := m[2][1]
	fmt.Println(m)
	fmt.Println(a, ok)
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "GOOD"
	a, ok = m[2][1]
	fmt.Println(m)
	fmt.Println(a, ok)
	delete(m, 2)
	fmt.Println("After delete key from map")
	fmt.Println(m)
	fmt.Println(a)
}

func iterationDemo1() {
	sm := make([]map[int]string, 5)
	//对slice迭代操作 下面对value的改变不能影响slice本身
	for _, v := range sm {
		//对map键值对的迭代操作
		v = make(map[int]string, 1)
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)
}

func iterationDemo2() {
	sm := make([]map[int]string, 5)
	//对slice迭代操作 使用slice的索引 对slice本身操作才能够有效
	for i := range sm {
		//对map键值对的迭代操作
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}

func sortDemo() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	//对slice []int 进行排序
	sort.Ints(s)
	fmt.Println(m)
	fmt.Println(s)
}

func homework_8() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m1)
	fmt.Println(m2)
}
