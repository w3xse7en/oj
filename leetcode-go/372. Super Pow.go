package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 2^0 % 1337 = 1
// 2^285 % 1337 = 1
// after 285 will be a loop
func superPowOwn(a int, b []int) int {
	var s strings.Builder
	for _, v := range b {
		s.WriteString(strconv.Itoa(v))
	}
	r, _ := strconv.ParseInt(s.String(), 10, 64)
	sum := 1.0
	for i := int64(0); i < r; i++ {
		sum *= math.Mod(float64(a), 1337)
		if sum > float64(1337) {
			sum = math.Mod(sum, 1337)
		}
	}
	return int(sum)
}
func main() {
	fmt.Println(superPow(2, []int{1, 0}))
	fmt.Println(superPow(2147483647, []int{2, 0, 0}))
	for i := 0; i <= 1000; i++ {
		fmt.Println(i, superPow(2, []int{i}))
	}
}

//(ab) mod m = ((a mod m)* (b mod m)) mod m
//(a mod m ) mod m = a mod m
//a mod m = (a mod m) mod m
//
//a**2 mod m = ((a mod m ) * (a mod m )) mod m
//a**2 mod m = (((a mod m) mod m ) * ((a mod m) mod m )) mod m
//a**2 mod m = ((((a mod m) mod m ) * ((a mod m) mod m )) mod m) mod m
//a**2 mod m = (((a mod m)**2) mod m) mod m
//a**2 mod m = ((a mod m)**2) mod m
//...
//then we can get
//a**10 mod m =((a mod m)**10) mod m

func superPow(a int, b []int) int {
	mod := 1337
	rt := 1
	for i, v := range b {
		if i == 0 {
			rt = (rt * powerAndMod(a, v, mod)) % mod
		} else {
			rt = ((powerAndMod(rt, 10, mod) % mod) * (powerAndMod(a, v, mod) % mod)) % mod
		}
	}
	return rt
}

func powerAndMod(a, p, mod int) int {
	r := 1
	for p > 0 {
		r = (r * a) % mod
		p--
	}
	return r
}
