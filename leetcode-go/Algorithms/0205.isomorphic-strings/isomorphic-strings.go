package problem0205

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	st := make(map[byte]byte)
	ts := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		sv, tv := s[i], t[i]
		_, sok := st[sv]
		_, tok := ts[tv]
		if sok == false && tok == false {
			st[sv] = tv
			ts[tv] = sv
		}
		if (sok == false && tok == true) || (sok == true && tok == false) {
			return false
		}
		if st[sv] != tv || ts[tv] != sv {
			return false
		}
	}
	return true
}
