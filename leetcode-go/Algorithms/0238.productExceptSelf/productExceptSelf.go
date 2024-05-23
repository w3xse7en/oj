package _238_productExceptSelf

func productExceptSelf(nums []int) []int {
	list := make([]int, len(nums))
	for i, _ := range list {
		list[i] = 1
	}
	var a, b = 1, 1
	for i, j := 1, len(nums)-2; i < len(nums); i++ {
		a *= nums[i-1]
		b *= nums[j+1]
		list[i] *= a
		list[j] *= b
		j--
	}
	return list
}
