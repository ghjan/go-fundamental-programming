package main

import "fmt"

func pointerDemo() {
	a := 1
	var p *int = &a
	fmt.Println(*p)
}

func ifDemo() {
	a := true
	if a, b, c := 1, 2, 3; a+b+c > 6 {
		fmt.Println("大于6")
	} else {
		fmt.Println("小于等于6")
		fmt.Println(a)
	}
	fmt.Println(a)
}

func forDemo1() {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)
	}
}

func forDemo2() {
	a := 1
	for a <= 3 {
		a++
		fmt.Println(a)
	}
}

func forDemo3() {
	a := 1
	for i := 0; i < 3; i++ {
		a++
	}
	fmt.Println(a)
}

//最平常的switch
func switchDemo1() {
	a := 1
	switch a {
	case 0:
		fmt.Printf("a=%d\n", a)
	case 1:
		fmt.Printf("a=%d\n", a)
	default:
		fmt.Printf("default, a=%d\n", a)
	}
}

//case使用条件表达式
func switchDemo2() {
	a := 1
	switch {
	case a >= 0:
		fmt.Printf("a>=0, a is %d\n", a)
		//继续检查下一个case
		fallthrough
	case a >= 1:
		fmt.Printf("a=%d\n", a)
	default:
		fmt.Printf("default, a=%d\n", a)
	}
}
func breakDemo() {
LABEL:
	for {
		for i := 0; i < 10; i++ {
			if i > 2 {
				//跳出LABEL同一层次的循环
				break LABEL
			} else {
				fmt.Println(i)
			}
		}
	}
}
func continueDemo() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			//继续LABEL同一层次的循环，所以就跳出了内部的无限循环
			continue LABEL
		}
	}
}

func gotoDemo() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			//goto
			goto LABEL
		}
	}
}

//switch带初始化语句
func switchDemo3() {
	switch a := 1; {
	case a >= 0:
		fmt.Printf("a>=0, a is %d\n", a)
		//继续检查下一个case
		fallthrough
	case a >= 1:
		fmt.Printf("a=%d\n", a)
	default:
		fmt.Printf("default, a=%d\n", a)
	}
}

func main() {
	fmt.Println("----pointerDemo-----")
	pointerDemo()
	fmt.Println("---ifDemo------")
	ifDemo()
	fmt.Println("-forDemo1--------")
	forDemo1()
	fmt.Println("----forDemo2-----")
	forDemo2()
	fmt.Println("----forDemo3-----")
	forDemo3()

	fmt.Println("-switchDemo1--------")
	switchDemo1()
	fmt.Println("----switchDemo2-----")
	switchDemo2()
	fmt.Println("----switchDemo3-----")
	switchDemo3()
	fmt.Println("----breakDemo-----")
	breakDemo()
	fmt.Println("----continueDemo-----")
	continueDemo()
	fmt.Println("----gotoDemo-----")
	gotoDemo()
}
