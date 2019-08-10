package main

import (
	"fmt"
	"sync"
	"time"
)

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	fmt.Println("Function in background, duration:", duration)
	wg.Done()
}
func dosomething2(millisecs time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	fmt.Println("Function in background, duration:", duration)
	panic("a")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go dosomething(200, &wg)
	go dosomething(400, &wg)
	go dosomething(150, &wg)
	go dosomething2(600, &wg)

	wg.Wait()
	fmt.Println("Done")
}
