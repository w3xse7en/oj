package main

import (
	"fmt"
)

// You are climbing a stair case. It takes n steps to reach to the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Note: Given n will be a positive integer.
//
// Example 1:
//
// Input: 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
// Example 2:
//
// Input: 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step

// -----------------------------------------------------------------------
// Mind
// 1
// 1
// sum 1

// 2
// 1 1
// 2
// sum 2

// 3
// 1 1 1
// 1 2
// 2 1
// sum = 1+2 = 3

// 4
// 1 1 1 1
// 1 1 2
// 1 2 1
// 2 1 1
// 2 2
// sum = 3+2 = 5

// 5
// 1 1 1 1 1
// 1 1 1 2
// 1 1 2 1
// 1 2 1 1
// 2 1 1 1
// 1 2 2
// 2 1 2
// 2 2 1
// sum = 5+3 = 8
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(climbStairs(1)) // 1
	fmt.Println(climbStairs(2)) // 2
	fmt.Println(climbStairs(3)) // 3
	fmt.Println(climbStairs(4)) // 5
	fmt.Println(climbStairs(5)) // 8
}
