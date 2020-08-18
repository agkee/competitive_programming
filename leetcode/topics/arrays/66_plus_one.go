package arrays

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	ret := make([]int, len(digits)+1)
	ret[0] = 1
	return ret
}
