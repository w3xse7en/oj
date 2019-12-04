package problem0125

func isPalindrome(s string) bool {
	var is, ij uint8
	for i, j := 0, len(s)-1; j > i && i < len(s) && j >= 0; {
		ui := checkAndToLower(s[i])
		if ui != 0 {
			is = ui
		} else {
			i++
		}
		uj := checkAndToLower(s[j])
		if uj != 0 {
			ij = uj
		} else {
			j--
		}
		if is != 0 && ij != 0 {
			if is == ij {
				ij = 0
				is = 0
				i++
				j--
			} else {
				return false
			}
		}
	}
	return true
}
func checkAndToLower(u uint8) uint8 {
	if (u >= 'A' && u <= 'Z') || (u >= 'a' && u <= 'z') || (u >= '0' && u <= '9') {
		if u >= 'A' && u <= 'Z' {
			u += 32
			return u
		}
		return u
	}
	return 0
}

func isPalindromeGoroutine(s string) bool {
	idx, jdx := make(chan int, 1), make(chan int, 1)
	is, js := make(chan uint8, 1), make(chan uint8, 1)
	go func() {
		for i := 0; i < len(s); i++ {
			if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
				str := s[i]
				if s[i] >= 'A' && s[i] <= 'Z' {
					str += 32
				}
				idx <- i
				is <- str
			}
		}
		close(idx)
	}()
	go func() {
		for j := len(s) - 1; j >= 0; j-- {
			if (s[j] >= 'A' && s[j] <= 'Z') || (s[j] >= 'a' && s[j] <= 'z') || (s[j] >= '0' && s[j] <= '9') {
				str := s[j]
				if s[j] >= 'A' && s[j] <= 'Z' {
					str += 32
				}
				jdx <- j
				js <- str
			}
		}
		close(jdx)
	}()

	for {
		i := <-idx
		j := <-jdx
		if i >= j {
			break
		}
		is := <-is
		js := <-js
		if is != js {
			return false
		}
	}

	return true
}
