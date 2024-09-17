package main

import "fmt"

// https://leetcode.cn/problems/minimum-additions-to-make-valid-string
func addMinimum(word string) int {
	res, l := 0, len(word)
	i := 0
	for i < l {
		if word[i] == 'a' {
			if i+1 < l && word[i+1] == 'b' {
				if i+2 < l && word[i+2] == 'c' {
					i += 3
				} else {
					res += 1
					i += 2
				}
			} else if i+1 < l && word[i+1] == 'c' {
				res += 1
				i += 2
			} else {
				res += 2
				i += 1
			}
		} else if word[i] == 'b' {
			res += 1
			if i+1 < l && word[i+1] == 'c' {
				i += 2
			} else {
				res += 1
				i += 1
			}
		} else if word[i] == 'c' {
			res += 2
			i += 1
		}
	}
	return res
}

func main071010() {
	fmt.Println(addMinimum("cba"))
}
