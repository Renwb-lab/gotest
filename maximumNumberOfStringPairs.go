package main

import "fmt"

// https://leetcode.cn/problems/find-maximum-number-of-string-pairs/description/
func maximumNumberOfStringPairs(words []string) int {
	reverse := func(s string) string {
		ans := []byte(s)
		for i, j := 0, len(ans)-1; i <= j; i, j = i+1, j-1 {
			ans[i], ans[j] = ans[j], ans[i]
		}
		return string(ans)
	}

	mp := map[string]int{}
	ans := 0
	for _, s := range words {
		rs := reverse(s)
		if _, ok := mp[rs]; ok {
			ans += 1
		} else {
			mp[s] = 1
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{
		"cd", "ac", "dc", "ca", "zz",
	}))
}
