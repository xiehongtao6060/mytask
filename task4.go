package main

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""
*/
func longestCommonPrefix(strs []string) string {
	// 边界情况：空数组
	if len(strs) == 0 {
		return ""
	}
	// 初始前缀为第一个字符串
	prefix := strs[0]
	// 遍历后续字符串，逐步缩小前缀
	for _, s := range strs[1:] {
		// 若当前字符串为空，直接返回空（无公共前缀）
		if len(s) == 0 {
			return ""
		}
		// 找到 prefix 与 s 的公共前缀长度
		i := 0
		for i < len(prefix) && i < len(s) && prefix[i] == s[i] {
			i++ // 字符匹配则继续后移
		}
		// 更新前缀为公共部分
		prefix = prefix[:i]
		// 若前缀已为空，提前退出
		if prefix == "" {
			return ""
		}
	}
	return prefix
}
