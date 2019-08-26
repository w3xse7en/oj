package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {
}

func live() People {
	var stu *Student
	return stu
}

// func main() {
// 	s := []int{0: 1, 1: 2, 2: 3, 3: 4,4:5,5:6}
// 	fmt.Printf("%p %p %p %p %p %p %v", s, s[1:], s[2:], s[3:],s[4:],s[], 1)
//
// }

func main() {
	c := make(chan interface{}, 1)
	c <- nil
	tc := <-c
	fmt.Println(tc)
}
