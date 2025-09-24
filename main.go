package main

import "fmt"

func task1() {
	var var1 = []int{1, 2, 3, 2, 1, 5, 5, 6, 6}
	var result = singleNumber(var1)
	fmt.Printf("hello world, result: %d", result)
}

func main() {
	task1()
}
