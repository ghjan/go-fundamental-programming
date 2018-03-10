package main

import (
	"fmt"
	"time"
)

func main() {
	timeFormatTrap()
}

//time.Format的坑
func timeFormatTrap() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("Mon Jan _2 15:04:06 2006"))
	//比上一个时间小 这就是 time.Format的坑
}
