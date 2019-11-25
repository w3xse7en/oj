package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Example 1:
//
// Input: a = "11", b = "1"
// Output: "100"
// Example 2:
//
// Input: a = "1010", b = "1011"
// Output: "10101"

func addBinaryBetter(a string, b string) string {
	carry, ans := 0, ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		if sum >= 2 {
			carry = 1
			sum -= 2
		} else {
			carry = 0
		}
		ans = strconv.Itoa(sum) + ans
	}
	return ans
}

func plusResult(a byte, b byte, plus bool) (byte, bool) {
	if a == '0' && b == '0' {
		if plus {
			return '1', false
		} else {
			return '0', false
		}
	}
	if (a == '1' && b == '0') || (a == '0' && b == '1') {
		if plus {
			return '0', true
		} else {
			return '1', false
		}
	}
	if a == '1' && b == '1' {
		if plus {
			return '1', true
		} else {
			return '0', true
		}
	}
	return 0, false
}
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	step := len(a) - len(b)
	var reverse, result strings.Builder
	var plus bool
	var byt byte
	for i := len(b) - 1; i >= 0; i-- {
		byt, plus = plusResult(a[i+step], b[i], plus)
		reverse.WriteByte(byt)
	}
	for i := step - 1; i >= 0; i-- {
		byt, plus = plusResult(a[i], '0', plus)
		reverse.WriteByte(byt)
	}
	if plus {
		reverse.WriteByte('1')
	}
	str := reverse.String()
	for i := reverse.Len() - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}
	return result.String()
}

func main() {
	fmt.Println(addBinary("0", "0"))       // 0
	fmt.Println(addBinary("0", "1"))       // 1
	fmt.Println(addBinary("1", "1"))       // 10
	fmt.Println(addBinary("10", "1"))      // 11
	fmt.Println(addBinary("11", "1"))      // 100
	fmt.Println(addBinary("11", "11"))     // 110
	fmt.Println(addBinary("1010", "1011")) // 10101
	// fmt.Println(fmt.Sprintf("%0b",0b11+0b11))
}
