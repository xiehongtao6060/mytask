package main

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。
*/
func plusOne(digits []int) []int {
	l := len(digits)

	plusOne := 0
	digits[l-1] = digits[l-1] + 1
	if digits[l-1] > 9 {
		digits[l-1] -= 10
		plusOne = 1
	} else {
		plusOne = 0
	}

	for i := l - 2; i >= 0; i-- {
		digits[i] = digits[i] + plusOne
		if digits[i] > 9 {
			digits[i] -= 10
			plusOne = 1
		} else {
			plusOne = 0
		}
	}

	if plusOne == 1 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}

}
