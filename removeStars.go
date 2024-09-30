package main

// https://leetcode.cn/problems/removing-stars-from-a-string/?envType=daily-question&envId=2024-09-21

func removeStars(s string) string {
	bs := []byte(s)
	st := make([]byte, 0)
	for _, b := range bs {
		if b == '*' {
			st = st[:len(st)-1]
		} else {
			st = append(st, b)
		}
	}
	return string(st)
}
