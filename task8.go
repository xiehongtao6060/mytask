package main

/*
两数之和

考察：数组遍历、map使用

题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
func twoSum(nums []int, target int) []int {
	// 使用单个map存储值到索引的映射（避免重复值覆盖）
	numMap := make(map[int]int)

	// 遍历数组，边遍历边查找互补值
	for i, num := range nums {
		complement := target - num // 计算互补值（目标值 - 当前元素）

		// 若互补值已存在于map中，说明找到两个数
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i} // 返回互补值索引和当前索引
		}

		// 将当前元素及其索引存入map（放在查找后，避免使用同一元素）
		numMap[num] = i
	}

	return nil // 题目保证有解，实际不会执行到此处
}
