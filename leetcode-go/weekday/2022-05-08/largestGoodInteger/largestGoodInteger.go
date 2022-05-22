package problem6056

func largestGoodInteger(num string) string {
	mp := map[string]bool{
		"000": false,
		"111": false,
		"222": false,
		"333": false,
		"444": false,
		"555": false,
		"666": false,
		"777": false,
		"888": false,
		"999": false,
	}
	for i := 0; i+3 <= len(num); i++ {
		t := num[i : i+3]
		_, ok := mp[t]
		if ok {
			mp[t] = true
		}
	}
	var max string
	for k, ok := range mp {
		if ok && k > max {
			max = k
		}
	}
	return max
}
