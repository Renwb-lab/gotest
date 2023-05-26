package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/russian-doll-envelopes/?company_slug=ubiquanthr

// 354. 俄罗斯套娃信封问题

// 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

// 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

// 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

// 注意：不允许旋转信封。

// 示例 1：

// 输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
// 输出：3
// 解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
// 示例 2：

// 输入：envelopes = [[1,1],[1,1],[1,1]]
// 输出：1

func maxEnvelopes1(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	// 使用动态规划
	helper := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i += 1 {
		helper[i] = 1
	}
	ret := helper[0]
	for i := 1; i < len(envelopes); i += 1 {
		temp := 1
		for j := i - 1; j >= 0; j -= 1 {
			if envelopes[j][1] < envelopes[i][1] {
				if helper[j]+1 > temp {
					temp = helper[j] + 1
				}
			}
		}
		helper[i] = temp
		if temp > ret {
			ret = temp
		}
	}

	return ret
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	helper := make([]int, len(envelopes))
	piles := -1
	for i := 0; i < len(envelopes); i += 1 {
		poker := envelopes[i][1]
		// 通过耐心排序
		// 具体为查找数组中最左边的数据
		left, right := 0, piles
		for left <= right {
			mid := left + (right-left)>>1
			if helper[mid] > poker {
				right = mid - 1
			} else if helper[mid] < poker {
				left = mid + 1
			} else if helper[mid] == poker {
				right = mid - 1
			}
		}
		if left > piles {
			piles += 1
		}
		helper[left] = poker
	}

	return piles + 1
}

func main93() {
	envelopes := [][]int{
		{46, 89},
		{50, 53},
		{52, 68},
		{72, 45},
		{77, 81},
	}
	fmt.Println(maxEnvelopes(envelopes))
	fmt.Println(convertIp("255.255.255.255"))
}

func convertIp(s string) int {
	ret, num, time := 0, 0, 0
	for i := len(s) - 1; i >= 0; i -= 1 {
		isDigit := s[i] >= '0' && s[i] <= '9'
		if isDigit {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] == '.' || i == 0 {
			for j := 0; j < time; j += 1 {
				num *= 256
			}
			ret += num
			num = 0
			time += 1
		}
	}
	return ret
}

func selectPair(arr []int, k int) [][]int {
	ret := make([][]int, 0)
	helper := make(map[int]int)
	for _, v := range arr {
		if n, ok := helper[k-v]; ok {
			for t := 0; t < n; t += 1 {
				ret = append(ret, []int{v, k - v})
			}
		} else {
			helper[v] = 1
		}
	}
	return ret
}
