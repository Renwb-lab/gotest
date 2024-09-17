package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	// f(0) = f(1) + f(2)
	l := len(s)
	var dp func(int) int
	m := make(map[int]int, 0)
	dp = func(idx int) int {
		if r, ok := m[idx]; ok {
			return r
		}
		if idx >= l {
			m[idx] = 1
			return 1
		}
		if s[idx] == '0' {
			return 0
		}
		t := 0
		if s[idx] >= '1' && s[idx] <= '9' {
			t += dp(idx + 1)
		}
		if idx+1 < l {
			n, err := strconv.Atoi(s[idx : idx+2])
			if err == nil {
				if n > 0 && n <= 26 {
					t += dp(idx + 2)
				}
			}
		}
		m[idx] = t
		return t
	}
	r := dp(0)
	return r
}

func numDecodings2(s string) int {
	if s[0] == '0' {
		return 0
	}
	l := len(s)
	helper := make([]int, l, l)
	if s[0] >= '1' && s[0] <= '9' {
		helper[0] = 1
	}
	for i := 1; i < l; i += 1 {
		helper[i] = 0
		if s[i] >= '1' && s[i] <= '9' {
			helper[i] += helper[i-1]
		}
		if n, err := strconv.Atoi(s[i-1 : i+1]); err == nil {
			if n >= 10 && n <= 26 {
				if i-2 < 0 {
					helper[i] += 1
				} else {
					helper[i] += helper[i-2]
				}
			}
		}
		if helper[i] == 0 {
			return 0
		}
	}
	return helper[l-1]
}

func numDecodings3(s string) int {
	if s[0] == '0' {
		return 0
	}
	l := len(s)
	a, b, c := 1, 0, 0
	if s[0] >= '1' && s[0] <= '9' {
		b = 1
	}
	if l == 1 {
		return b
	}
	for i := 1; i < l; i += 1 {
		c = 0
		if s[i] >= '1' && s[i] <= '9' {
			c += b
		}
		if n, err := strconv.Atoi(s[i-1 : i+1]); err == nil {
			if n >= 10 && n <= 26 {
				c += a
			}
		}
		if c == 0 {
			return 0
		}
		a, b = b, c
	}
	return c
}

func main288() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings2("2101"))
	fmt.Println(numDecodings3("2101"))
}
