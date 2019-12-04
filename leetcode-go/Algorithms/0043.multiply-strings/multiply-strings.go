package problem0043

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	n1 := reverseToIntList(num1)
	n2 := reverseToIntList(num2)
	det := [][]int{}
	resultIntList := []int{}
	for k, v := range n2 {
		list := multiplyByIntList(n1, v)
		if len(list) != 0 {
			det = append(det, append(make([]int, k), list...))
		}
	}

	plus := 0
	length := 0
	if len(det) != 0 {
		length = len(det[len(det)-1])
	}
	for i := 0; i < length; i++ {
		deti := 0
		for _, v := range det {
			vi := 0
			if i < len(v) {
				vi = v[i]
			}
			deti += vi
		}
		deti += plus

		plus = deti / 10
		deti = deti % 10
		resultIntList = append(resultIntList, deti)
	}
	if plus != 0 {
		resultIntList = append(resultIntList, plus)
	}
	if len(resultIntList) == 0 {
		return "0"
	}
	var result strings.Builder
	for i := len(resultIntList) - 1; i >= 0; i-- {
		result.WriteString(strconv.Itoa(resultIntList[i]))
	}
	return result.String()
}

func multiplyByIntList(list []int, n int) []int {
	result := []int{}
	if n == 0 {
		return result
	}
	plus := 0
	for _, v := range list {
		vi := (v * n) + plus
		plus = vi / 10
		vi = vi % 10
		result = append(result, vi)
	}
	if plus != 0 {
		result = append(result, plus)
	}
	return result
}

func reverseToIntList(str string) []int {
	result := []int{}
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, int(str[i]-'0'))
	}
	return result
}
