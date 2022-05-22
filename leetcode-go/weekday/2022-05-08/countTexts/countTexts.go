package countTexts

import "fmt"

func countTexts(pressedKeys string) int {
	d(5)
	//three, mt, four, mf := getList(pressedKeys)
	//fmt.Println("mt", mt, "three", three)
	//fmt.Println("mf", mf, "four", four)
	fmt.Println("---")
	return 0
}

func d(i int) int {
	if i == 0 {
		return 1
	}
	r := d(i-1) + d(i-2)
	fmt.Println(i, r)
	return r
}

func getList(pressedKeys string) ([]int, int, []int, int) {
	three, maxThree := []int{}, 0

	four, maxFour := []int{}, 0
	var t rune
	cnt := 0
	if len(pressedKeys) != 0 {
		t = rune(pressedKeys[0])
	}
	for _, v := range pressedKeys {
		if t == v {
			cnt++
			continue
		}
		if t == '7' || t == '9' {
			four = append(four, cnt)
		} else {
			three = append(three, cnt)
		}
		t = v
		cnt = 1
	}
	if t == '7' || t == '9' {
		four = append(four, cnt)
		if cnt > maxFour {
			maxFour = cnt
		}
	} else {
		three = append(three, cnt)
		if cnt > maxThree {
			maxThree = cnt
		}
	}
	return three, maxThree, four, maxFour
}
