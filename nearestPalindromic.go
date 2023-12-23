package main

import (
	"fmt"
	"strconv"
)

func nearestPalindromic1(n string) string {
	dec := func(n string) string {
		l := len(n)
		res := make([]byte, l, l)
		hasDec := false
		for i := l - 1; i >= 0; i -= 1 {
			if hasDec {
				res[i] = n[i]
			} else {
				if n[i] == '0' {
					res[i] = '9'
				} else {
					res[i] = n[i] - 1
					hasDec = true
				}
			}
		}
		i := 0
		for ; i < l && res[i] == '0'; i += 1 {
		}
		if i == l {
			return "0"
		}
		return string(res[i:])
	}
	add := func(n string) string {
		n = fmt.Sprintf("0%s", n)
		l := len(n)
		res := make([]byte, l, l)
		hasAdd := false
		for i := l - 1; i >= 0; i -= 1 {
			if hasAdd {
				res[i] = n[i]
			} else {
				if n[i] == '9' {
					res[i] = '0'
				} else {
					res[i] = n[i] + 1
					hasAdd = true
				}
			}
		}
		i := 0
		for ; i < l && res[i] == '0'; i += 1 {
		}
		return string(res[i:])
	}
	palindromic := func(n string) bool {
		l := len(n)
		for a, b := 0, l-1; a <= b; a, b = a+1, b-1 {
			if n[a] != n[b] {
				return false
			}
		}
		return true
	}
	n1, n2 := n, n
	for {
		n1 = dec(n1)
		if palindromic(n1) {
			return n1
		}
		n2 = add(n2)
		if palindromic(n2) {
			return n2
		}
	}
}

func nearestPalindromic(n string) string {
	abs := func(n int) int {
		if n < 0 {
			return -n
		} else {
			return n
		}
	}
	reverse := func(x string) string {
		l := len(x)
		res := make([]byte, l, l)
		for i := l - 1; i >= 0; i -= 1 {
			res[i] = x[l-1-i]
		}
		return string(res)
	}
	helper := func(x string, flag bool) string {
		if flag {
			return fmt.Sprintf("%s%s", x, reverse(x[:len(x)-1]))
		}
		return fmt.Sprintf("%s%s", x, reverse(x))

	}

	l := len(n)
	candidates := make([]string, 0)
	// 1001
	c1 := make([]byte, l+1, l+1)
	c1[0], c1[l] = '1', '1'
	for i := 1; i < l; i += 1 {
		c1[i] = '0'
	}
	candidates = append(candidates, string(c1))
	// 999
	c2 := make([]byte, l-1, l-1)
	for i := 0; i < l-1; i += 1 {
		c2[i] = '9'
	}
	candidates = append(candidates, string(c2), "0")
	prefix := n[:(l+1)/2]
	prefixNum, _ := strconv.Atoi(prefix)
	addfix, decfix := strconv.Itoa(prefixNum+1), strconv.Itoa(prefixNum-1)
	for _, x := range []string{decfix, prefix, addfix} {
		candidates = append(candidates, helper(x, l%2 == 1))
	}

	intR, res := -1, ""
	intN, _ := strconv.Atoi(n)
	for _, c := range candidates {
		if intC, err := strconv.Atoi(c); err == nil {
			if intC != intN {
				if intR == -1 ||
					abs(intC-intN) < abs(intR-intN) ||
					(abs(intC-intN) == abs(intR-intN) && intC < intR) {
					intR = intC
					res = c
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(nearestPalindromic1("1213"))
	fmt.Println(nearestPalindromic("1"))
}
