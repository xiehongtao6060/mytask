package main

/*
回文数

考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10

	}

	return x == reversedHalf || x == reversedHalf/10
}
