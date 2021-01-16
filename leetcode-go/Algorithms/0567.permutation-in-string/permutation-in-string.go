package problem0567

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	i, j := 0, 0

	needMp := map[byte]int{}
	for k := range s1 {
		needMp[s1[k]]++
	}
	match := 0
	winMp := map[byte]int{}
	for ; j < len(s1); j++ {
		v := s2[j]
		winMp[v]++
		if winMp[v] == needMp[v] {
			match++
		}
	}
	if match == len(needMp) {
		return true
	}
	for j < len(s2) {
		pop := s2[i]
		push := s2[j]
		if winMp[pop] == needMp[pop] {
			match--
		}
		winMp[pop]--
		winMp[push]++
		if winMp[push] == needMp[push] {
			match++
		}
		if match == len(needMp) {
			return true
		}
		i++
		j++
	}
	return false
}
