package problem0005

//
func longestPalindrome(s string) string {
	max, result := 0, ""
	for k := range s {
		i := 0
		for {
			if k+1-i > max {
				is := true
				ti, tj := i, k
				for ti < tj {
					if s[ti] != s[tj] {
						is = false
						break
					}
					ti++
					tj--
				}
				if is {
					s2 := s[i : k+1]
					max = len(s2)
					result = s2
				}
				i++
			} else {
				break
			}
		}
	}
	return result
}
