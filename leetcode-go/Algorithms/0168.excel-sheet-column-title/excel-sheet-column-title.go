package problem0168

// 注意下标从0开始
// A的实际映射是0 而不是1
// Z的映射是25 而不是26
func convertToTitle(n int) string {
	var result []byte
	for ; n != 0; n /= 26 {
		n--
		result = append([]byte{byte('A' + n%26)}, result...)
	}
	return string(result)
}
