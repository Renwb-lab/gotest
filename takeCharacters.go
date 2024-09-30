package main

// https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right/
func takeCharacters(s string, k int) int {
	charToNum := make([]int, 3)
	// 统计每个字符出现的次数
	for _, c := range s {
		charToNum[c-'a'] += 1
	}
	ans := len(s)
	// 如果每个字符出现的次数都大于k的话，则表示最多可以使用len(s)
	if charToNum[0] >= k && charToNum[1] >= k && charToNum[2] >= k {
		ans = min(ans, len(s))
	} else {
		return -1
	}
	// 对于扣减后剩下的字符串使用滑动窗口计算
	left, right := 0, 0
	for ; right < len(s); right += 1 {
		// 如果right移动，则表明对应的字符数量-1
		charToNum[s[right]-'a'] -= 1
		// 如果不能满足要求，则需要有移动left
		// 不能满足的要求为a,b,c的数量不少于k
		for left < right && (charToNum[0] < k || charToNum[1] < k || charToNum[2] < k) {
			// 右移left, 并且对应的字符数量+1
			charToNum[s[left]-'a'] += 1
			left += 1
		}
		// 如果右移left后满足题目要求，则更新结果
		if charToNum[0] >= k || charToNum[1] >= k || charToNum[2] >= k {
			ans = min(ans, len(s)-(right-left+1))
		}
	}
	return ans
}
