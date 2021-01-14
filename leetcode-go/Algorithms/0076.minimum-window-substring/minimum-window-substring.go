package problem0076

func minWindow(s string, t string) string {
	win := make([]rune, 0, len(s))
	tMp, wMp := map[rune]int{}, map[rune]int{}
	for _, v := range t {
		tMp[v]++
	}
	minLen := len(s) + 1
	tMpLen := len(tMp)
	match := 0
	var result string
	for _, v := range s {
		win = append(win, v)
		//fmt.Println(1,string(win))
		wMp[v]++
		if _, ok := tMp[v]; ok && wMp[v] == tMp[v] {
			match++
		}
		if match != tMpLen {
			continue
		}
		winLen := len(win)
		if minLen > winLen {
			minLen = winLen
			result = string(win)
		}
		for {
			pop := win[0]
			wMp[pop]--
			win = win[1:]
			//fmt.Println(2,string(win))
			if _, ok := tMp[pop]; ok && wMp[pop] < tMp[pop] {
				match--
			}
			if match != tMpLen {
				break
			}
			winLen := len(win)
			if minLen > winLen {
				minLen = winLen
				result = string(win)
			}
		}
	}
	return result
}
