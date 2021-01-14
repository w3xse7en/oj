package problem0567

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	match := 0
	s1Mp, winMp := map[byte]int{}, map[byte]int{}
	for i := range s1 {
		s1Mp[s1[i]]++
	}

	for i := 0; i < len(s1); i++ {
		v := s2[i]
		winMp[v]++
		if _, ok := s1Mp[v]; ok && winMp[v] == s1Mp[v] {
			match++
		}
	}

	i, j := 0, len(s1)
	for {
		if match == len(s1Mp) {
			return true
		}
		if j >= len(s2) {
			break
		}
		pop := s2[i]
		winMp[pop]--
		if _, ok := s1Mp[pop]; ok && winMp[pop] < s1Mp[pop] {
			match--
		}
		push := s2[j]
		winMp[push]++
		if _, ok := s1Mp[push]; ok && winMp[push] == s1Mp[push] {
			match++
		}

		i++
		j++
	}
	return false
}
