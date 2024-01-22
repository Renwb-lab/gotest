package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	f := func(words []string) (ans map[string]int) {
		ans = map[string]int{}
		for _, s := range words {
			ans[s] += 1
		}
		return
	}
	m1, m2 := f(words1), f(words2)
	ans := 0
	for k1, v1 := range m1 {
		if v1 == 1 {
			if v2, ok := m2[k1]; ok && v2 == 1 {
				ans += 1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countWords(
		[]string{"leetcode", "is", "amazing", "as", "is"},
		[]string{"amazing", "leetcode", "is"}))
}
