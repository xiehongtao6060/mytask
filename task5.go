package main

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。
*/
func plusOne(digits []int) []int {
	carry := 1 // 初始进位为1（即要加的1）

	// 从末尾（最低位）开始处理所有位的进位逻辑
	for i := len(digits) - 1; i >= 0 && carry > 0; i-- {
		sum := digits[i] + carry
		if sum == 10 {
			digits[i] = 0 // 当前位变为0，进位继续
			carry = 1
		} else {
			digits[i] = sum // 当前位直接赋值，进位消化
			carry = 0
		}
	}

	// 若所有位处理完仍有进位（全9情况），在头部插入1
	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits

}
