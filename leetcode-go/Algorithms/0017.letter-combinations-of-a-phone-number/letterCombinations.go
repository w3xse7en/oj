package problem0017

var b = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	list := b[digits[0]]
	for i := 1; i < len(digits); i++ {
		nList := []string{}
		for _, lv := range list {
			for _, bv := range b[digits[i]] {
				nList = append(nList, lv+bv)
			}
		}
		list = nList
	}
	return list
}
