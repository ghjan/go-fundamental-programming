package main

import "fmt"

func main() {
	homework_7()
}
func homework_7() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:]
	fmt.Println(s1)
	fmt.Println(s2)
}
