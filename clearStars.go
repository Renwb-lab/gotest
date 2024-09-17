package main

import "fmt"

func clearStars(s string) string {
	l := len(s)
	h := make([][]int, 26)
	for i := 0; i < 26; i += 1 {
		h[i] = make([]int, 0)
	}
	deleteFlag := make([]bool, len(s))

	for i, c := range s {
		// fmt.Println(h)
		if c != '*' {
			h[c-'a'] = append(h[c-'a'], i)
		} else {
			deleteFlag[i] = true
			for i, l := range h {
				if len(l) > 0 {
					idx := l[len(l)-1]
					h[i] = l[:len(l)-1] //注意这里
					deleteFlag[idx] = true
					break
				}
			}
		}
	}

	ans := []byte{}
	for i := 0; i < l; i += 1 {
		if !deleteFlag[i] {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

func main071015() {
	fmt.Println(clearStars("d*o*a"))
}
