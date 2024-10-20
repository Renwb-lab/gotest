package main

import "sort"

// https://leetcode.cn/problems/check-if-a-string-can-break-another-string/
func checkIfCanBreak(s1 string, s2 string) bool {
	bs1, bs2 := []byte(s1), []byte(s2)
	sort.Slice(bs1, func(i, j int) bool { return bs1[i] < bs1[j] })
	sort.Slice(bs2, func(i, j int) bool { return bs2[i] < bs2[j] })
	compare := func(bs1, bs2 []byte) bool {
		for i := range bs1 {
			if bs1[i] > bs2[i] {
				return false
			}
		}
		return true
	}

	return compare(bs1, bs2) || compare(bs2, bs1)
}

func checkIfCanBreak2(s1 string, s2 string) bool {
	compare := func(s1, s2 string) bool {
		m := make([]int, 26)
		for _, c := range s1 {
			m[c-'a'] += 1
		}
		for _, c := range s2 {
			canFind := false
			for i := c - 'a'; i < 26; i += 1 {
				if m[i] > 0 {
					m[i] -= 1
					canFind = true
					break
				}
			}
			if !canFind {
				return false
			}
		}
		return true
	}

	return compare(s1, s2) || compare(s2, s1)
}
