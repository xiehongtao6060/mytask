package main

import "fmt"

func task1() {
	var var1 = []int{1, 2, 3, 2, 1, 5, 5, 6, 6}
	var result = singleNumber(var1)
	fmt.Printf("hello world, result: %d", result)
}

func task2() {
	x := 1221
	if isPalindrome(x) {
		fmt.Printf("hello world, task2 is true")
	} else {
		fmt.Printf("hello world, task2 is false")
	}

}

func task3() {
	s := "()"
	result := isValid(s)

	fmt.Printf("hello world, result: %t", result)

}

func task4() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Printf("hello world, result: %s", result)
}

func task5() {
	digits := []int{9}
	digits = plusOne(digits)
	fmt.Println("result:", digits)
}

func task6() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("input:", nums)
	k := removeDuplicates(nums)
	fmt.Println("result:", k, nums)

}

func task7() {
	intervals := [][]int{{1, 3}, {1, 4}}
	intervals = merge(intervals)
	fmt.Println("intervals:", intervals)

}
func main() {
	task7()
}
