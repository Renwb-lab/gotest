package main

import "fmt"

// https://leetcode.cn/problems/permutation-sequence/submissions/515710268/

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i += 1 {
		factorial[i] = factorial[i-1] * i
	}
	used := make([]bool, n+1, n+1)
	getIndex := func(idx int) int {
		cnt := factorial[n-1-idx]
		for i := 1; i <= n; i += 1 {
			if used[i] {
				continue
			}
			if cnt < k {
				k -= cnt
				continue
			}
			used[i] = true
			return i
		}
		return 0
	}

	res := make([]byte, 0)
	for i := 0; i < n; i += 1 {
		idx := getIndex(i)
		res = append(res, byte('0'+idx))
	}
	return string(res)
}

func getPermutationV1(n int, k int) string {
	factorial := make([]int, n+1, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i += 1 {
		factorial[i] = factorial[i-1] * i
	}
	used := make([]bool, n+1, n+1)
	res := make([]byte, 0)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			return
		}
		cnt := factorial[n-1-idx]
		for i := 1; i <= n; i += 1 {
			if used[i] {
				continue
			}
			if cnt < k {
				k -= cnt
				continue
			}
			res = append(res, byte(i+'0'))
			used[i] = true
			dfs(idx + 1)
		}
	}
	dfs(0)
	return string(res)
}

func main() {
	fmt.Println(getPermutation(4, 9))
}
