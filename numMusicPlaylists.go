package main

import "fmt"

// https://leetcode.cn/problems/number-of-music-playlists/description/?company_slug=ubiquanthr

// 920. 播放列表的数量

// 你的音乐播放器里有 N 首不同的歌，在旅途中，你的旅伴想要听 L 首歌（不一定不同，即，允许歌曲重复）。请你为她按如下规则创建一个播放列表：

// 每首歌至少播放一次。
// 一首歌只有在其他 K 首歌播放完之后才能再次播放。
// 返回可以满足要求的播放列表的数量。由于答案可能非常大，请返回它模 10^9 + 7 的结果。

// 示例 1：
//
// 输入：N = 3, L = 3, K = 1
// 输出：6
// 解释：有 6 种可能的播放列表。[1, 2, 3]，[1, 3, 2]，[2, 1, 3]，[2, 3, 1]，[3, 1, 2]，[3, 2, 1].
// 示例 2：
//
// 输入：N = 2, L = 3, K = 0
// 输出：6
// 解释：有 6 种可能的播放列表。[1, 1, 2]，[1, 2, 1]，[2, 1, 1]，[2, 2, 1]，[2, 1, 2]，[1, 2, 2]
// 示例 3：
//
// 输入：N = 2, L = 3, K = 1
// 输出：2
// 解释：有 2 种可能的播放列表。[1, 2, 1]，[2, 1, 2]
//
//
// 提示：
//
// 0 <= K < N <= L <= 100

func numMusicPlaylists(n int, l int, k int) int {
	// 令 dp[i][j] 为播放列表长度为 i 包含恰好 j 首不同歌曲的数量。
	// 我们可以播放没有播放过的歌也可以是播放过的。如果未播放过的，那么就是 dp[i-1][j-1] * (N-j) 种选择方法。
	// 如果不是，那么就是选择之前的一首歌，dp[i-1][j] * max(j-K, 0)（j 首歌，最近的 K 首不可以播放）。

	Mod := 1000000007
	helper := make([][]int, l+1)
	for i := 0; i <= l; i += 1 {
		helper[i] = make([]int, n+1)
	}
	helper[0][0] = 1

	for i := 1; i <= l; i += 1 {
		for j := 1; j <= n; j += 1 {
			helper[i][j] += helper[i-1][j-1] * (n - j + 1)
			if j-k > 0 {
				helper[i][j] += helper[i-1][j] * (j - k)
			}
			helper[i][j] %= Mod
		}
	}

	return helper[l][n]
}

func main39() {
	n, goal, k := 3, 3, 1
	fmt.Println(numMusicPlaylists(n, goal, k))
}
