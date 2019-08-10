package main

import (
	"fmt"
	"time"
)

func main() {
	emptychan1 := make(chan int)
	go func() {
		emptychan1 <- 1
	}()
	<-emptychan1
	time.Sleep(time.Second)
	// close(emptychan1)
	select {
	case a, ok := <-emptychan1:
		fmt.Println(a, ok)
	}

}
