package main

import (
	"fmt"
	"time"
)

// 两个协程交替打印
func main() {
	toStr := make(chan int)
	toNumb := make(chan int)
	go func() {
		for i := 1; ; i++ {
			select {
			case r := <-toNumb:
				fmt.Println(i)
				i++
				fmt.Println(i)
				toStr <- i
				if r > 26 {
					return
				}
			}
		}
	}()
	go func() {
		for i := 0; ; i++ {
			select {
			case r := <-toStr:
				fmt.Println(string(byte('a' + i)))
				i++
				fmt.Println(string(byte('a' + i)))
				toNumb <- i
				if r > 26 {
					return
				}
			}
		}
	}()
	toNumb <- 0
	time.Sleep(time.Second)
}
