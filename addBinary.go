package main

import "fmt"

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	s := make([]byte, la+lb+1, la+lb+1)
	inc := 0
	x, y := la-1, lb-1
	loc := la + lb
	for x >= 0 && y >= 0 {
		t := int(a[x]-'0') + int(b[y]-'0') + inc
		if t == 2 {
			s[loc] = '0'
			inc = 1
		} else if t == 3 {
			s[loc] = '1'
			inc = 1
		} else {
			s[loc] = byte('0' + t)
			inc = 0
		}
		x, y, loc = x-1, y-1, loc-1
	}
	for x >= 0 {
		t := int(a[x]-'0') + inc
		if t == 2 {
			s[loc] = '0'
			inc = 1
		} else {
			s[loc] = byte('0' + t)
			inc = 0
		}
		x, loc = x-1, loc-1
	}
	for y >= 0 {
		t := int(b[y]-'0') + inc
		if t == 2 {
			s[loc] = '0'
			inc = 1
		} else {
			s[loc] = byte('0' + t)
			inc = 0
		}
		y, loc = y-1, loc-1
	}
	if inc == 1 {
		s[loc] = '1'
		return string(s[loc:])
	}
	return string(s[loc+1:])
}

func main0710() {
	fmt.Println(addBinary("0", "0"))
}
