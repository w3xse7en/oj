package main

import (
	"fmt"
	"sync"
	"time"
)

type Ban struct {
	visitIPs sync.Map
}

func NewBan() *Ban {
	var a sync.Map
	return &Ban{visitIPs: a}
}
func (o *Ban) visit(ip string) bool {
	v, loaded := o.visitIPs.LoadOrStore(ip, time.Now())
	if loaded {
		if time.Now().After(v.(time.Time).Add(time.Second * 1)) {
			o.visitIPs.Store(ip, time.Now())
			return true
		} else {
			return false
		}
	}
	return true
}
func main() {
	success := 0
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if ban.visit(ip) {
					success++
				}
			}(j)
		}
	}
	fmt.Println("success:", success)
}
