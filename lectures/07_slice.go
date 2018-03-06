package main

import "fmt"

func main() {
	fmt.Println("----sliceDemo---")
	sliceDemo()
	fmt.Println("----resliceDemo---")
	resliceDemo()
	fmt.Println("----appendDemo---")
	appendDemo()
	fmt.Println("----sliceDemo2---")
	sliceDemo2()
	fmt.Println("----appendDemo2---")
	appendDemo2()
	fmt.Println("----copyDemo---")
	copyDemo()
}
func sliceDemo() {
	//slice 中括号里面没有数字 也没有...
	var s1 []int
	fmt.Println(s1)
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//所以从0开始 9是最后一个
	s2 := a[9]
	fmt.Println(s2)
	s3 := a[5:]
	fmt.Println(s3)
	//使用make 比较正式 完整的属性
	s4 := make([]int, 3, 10)
	fmt.Println(s4)

	a5 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa5 := a5[2:5]
	fmt.Println(string(sa5))
	sb5 := a5[3:5]
	fmt.Println(string(sb5))
}

//从一个slice中获取一个新的slice
//reslice中用到的索引，以被reslice的slice为准
func resliceDemo() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	fmt.Println(string(sa))
	sb := sa[0:]
	fmt.Println(string(sb))
	sc := sa[3:5]
	fmt.Println(string(sc))
}

func appendDemo() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1)
	//加3个元素
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	//追加以后，如果数量大于原来的容量 就会扩充容量 slice的地址也会变化
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)

}

//-一般使用make()创建-如果多个slice指向相同底层数组，其中一个的值改变会影响全部
func sliceDemo2() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	s1[0] = 9
	fmt.Println(s1, s2)
}

func appendDemo2() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	//s2扩容以后使用了新地址 这个时候不再和s1共享地址了，s1中的内容改变也不会影响到s2
	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	s1[0] = 9
	fmt.Println(s1, s2)
}

func copyDemo() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	//dst, src
	copy(s2, s1)
	fmt.Println(s2)
}
