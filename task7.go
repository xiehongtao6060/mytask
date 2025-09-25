package main

import (
	"sort"
)

/*
56. 合并区间：以数组 intervals 表示若干个区间的集合，
其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func merge(intervals [][]int) [][]int {
	// 边界条件：空区间直接返回
	if len(intervals) == 0 {
		return nil
	}

	// 按区间起始位置排序（关键步骤：确保重叠区间相邻）
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片，放入第一个区间
	result := [][]int{intervals[0]}

	// 遍历剩余区间，合并重叠部分
	for _, current := range intervals[1:] {
		// 获取结果中最后一个已合并的区间
		last := result[len(result)-1]

		if current[0] <= last[1] {
			// 重叠：更新最后区间的结束值为两者最大值
			last[1] = max(last[1], current[1])
		} else {
			// 不重叠：直接加入结果
			result = append(result, current)
		}
	}

	return result
}

// 辅助函数：返回两个数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
