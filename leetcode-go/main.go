package main

import "fmt"

//func main() {
//	fmt.Println('a' - 'A')
//	fmt.Println("hello world")
//}

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("%v", p)
}

func main() {
	p := &People{}
	p.String()
}
