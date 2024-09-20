package main

import "strconv"

// https://leetcode.cn/problems/count-special-integers/
func countSpecialNumbers(n int) int {
	nStr := strconv.Itoa(n)
	// 等于k位的数量
	memo := make(map[int]int)
	var dp func(int, bool) int
	dp = func(mask int, prefixSmaller bool) int {
		if countOnes(mask) == len(nStr) {
			return 1
		}
		key := mask * 2
		if prefixSmaller {
			key += 1
		}
		if res, ok := memo[key]; ok {
			return res
		}
		res, lowerBound := 0, 0
		if mask == 0 {
			lowerBound = 1
		}
		upperBound := 9
		if !prefixSmaller {
			upperBound = int(nStr[countOnes(mask)] - '0')
		}
		for i := lowerBound; i <= upperBound; i += 1 {
			if (mask>>i)&1 == 0 {
				res += dp(mask|(1<<i), prefixSmaller || i < upperBound)
			}
		}
		memo[key] = res
		return res
	}
	// 小于k位的数量
	// 第一位:9 (1-9)
	// 第二位:9 (扣除第一位的数字)
	// 第二位:8 (扣除第一位、第二位的数字)
	res, prod := 0, 9
	for i := 0; i < len(nStr)-1; i++ {
		res += prod
		prod *= 9 - i
	}
	res += dp(0, false)
	return res
}

func countOnes(x int) int {
	count := 0
	for x > 0 {
		count += 1
		x &= x - 1
	}
	return count
}
