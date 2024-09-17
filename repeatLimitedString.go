package main

import "fmt"

func repeatLimitedString(s string, repeatLimit int) string {
	N := 26
	arr := make([]int, N)
	for _, c := range s {
		arr[c-'a'] += 1
	}
	var ret []byte
	count := 0
	for i, j := N-1, N-2; i >= 0 && j >= 0; {
		switch {
		case arr[i] == 0:
			count, i = 0, i-1
		case count < repeatLimit:
			arr[i] -= 1
			ret = append(ret, 'a'+byte(i))
			count += 1
		case arr[j] == 0 || j >= i: // // 当前字符已经超过限制，查找可填入的其他字符
			j -= 1
		default:
			arr[j] -= 1
			ret = append(ret, 'a'+byte(j))
			count = 0
		}
	}
	return string(ret)
}

func repeatLimitedString2(s string, repeatLimit int) string {
	arr := make([]int, 26)
	for _, c := range s {
		arr[c-'a'] += 1
	}
	charArr, numArr := make([]byte, 0), make([]int, 0)
	for i, n := range arr {
		if n > 0 {
			charArr = append(charArr, byte(rune('a'+i)))
			numArr = append(numArr, n)
		}
	}
	res := []byte{}
	idx := len(charArr) - 1
	count := 0
	next := func() bool {
		nextIdx := idx - 1
		for nextIdx >= 0 && numArr[nextIdx] == 0 {
			nextIdx -= 1
		}
		if nextIdx >= 0 && numArr[nextIdx] > 0 {
			res = append(res, charArr[nextIdx])
			numArr[nextIdx] -= 1
			count = 1
			return true
		}
		return false
	}
	for idx >= 0 {
		for idx >= 0 && numArr[idx] == 0 {
			idx -= 1
		}
		if idx < 0 {
			break
		}
		if len(res) == 0 {
			res = append(res, charArr[idx])
			numArr[idx] -= 1
			count = 1
			continue
		}
		if res[len(res)-1] != charArr[idx] {
			res = append(res, charArr[idx])
			numArr[idx] -= 1
			count = 1
			continue
		}
		if count < repeatLimit {
			res = append(res, charArr[idx])
			numArr[idx] -= 1
			count += 1
			continue
		}
		if !next() {
			break
		}
	}
	return string(res)
}

func main298() {
	fmt.Println(repeatLimitedString("aababab", 2))
}
