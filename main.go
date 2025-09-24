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

func main() {
	task4()
}
