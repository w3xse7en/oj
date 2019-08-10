package main

import (
	"container/list"
	"fmt"
	"reflect"
)

type State struct{}

func testnil1(a, b interface{}) bool {
	return a == b
}

func testnil2(a *State, b interface{}) bool {
	return a == b
}

func testnil3(a interface{}) bool {
	return a == nil
}

func testnil4(a *State) bool {
	list.New()
	return a == nil
}

func testnil5(a interface{}) bool {
	v := reflect.ValueOf(a)
	return !v.IsValid() || v.IsNil()
}

func testnil6(a, b *State) bool {
	return a == b
}

func main() {
	var a *State
	fmt.Println(testnil1(a, nil))
	fmt.Println(testnil2(a, nil))
	fmt.Println(testnil3(a))
	fmt.Println(testnil4(a))
	fmt.Println(testnil5(a))
	fmt.Println(testnil3(nil))
	fmt.Println(testnil4(nil))
	fmt.Println(testnil6(a, nil))
}
