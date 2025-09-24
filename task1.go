package main

func singleNumber(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	for num, cnt := range counts {
		if cnt == 1 {
			return num
		}
	}

	return -1
}
