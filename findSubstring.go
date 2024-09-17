package main

import "fmt"

func findSubstring(s string, words []string) []int {
	need := map[string]int{}
	for i := range words {
		need[words[i]] += 1
	}

	l := len(words[0])
	res := make([]int, 0)
	for i := 0; i < len(s); i += 1 {
		valid := 0
		m := map[string]int{}
		for j := i; j < len(s); j += l {
			if j+l > len(s) {
				j += 1
				break
			}
			w := s[j : j+l]
			n, ok := need[w]
			if !ok {
				break
			}

			m[w] += 1
			if m[w] > n {
				break
			}
			if m[w] == n {
				valid += 1
			}
			if valid == len(need) {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func main071045() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}
