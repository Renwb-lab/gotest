package main

// https://leetcode.cn/problems/maximize-win-from-two-segments/?envType=daily-question&envId=2024-09-21

func maximizeWin(prizePositions []int, k int) int {
	// dp[x] 表示以x-1为结束点第一段获取的最大长度
	// ans = right - left + 1 + dp[left]
	n := len(prizePositions)
	dp := make([]int, n+1)
	ans := 0
	for left, right := 0, 0; right < n; right += 1 {
		// 以right结尾时，最大可以获取的数量
		for prizePositions[right]-prizePositions[left] > k {
			left += 1
		}
		// right-left+1 表示第二个以right为右端点left为左端点时获取的数量
		// dp[left]表示第一个以left-1为右端点时获取的数量
		ans = max(ans, right-left+1+dp[left])
		dp[right+1] = max(dp[right], right-left+1)
	}
	return ans
}
