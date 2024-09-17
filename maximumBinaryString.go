package main

import (
	"fmt"
	"strings"
)

func maximumBinaryString(binary string) string {
	n := len(binary)
	i := strings.Index(binary, "0")
	if i < 0 {
		return binary
	}
	zeros := strings.Count(binary, "0")
	s := []rune(strings.Repeat("1", n))
	s[i+zeros-1] = '0'
	return string(s)
}

func maximumBinaryStringV1(binary string) string {
	n := len(binary)
	s := []rune(binary)
	j := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			for j <= i || (j < n && s[j] == '1') {
				j++
			}
			if j < n {
				s[j] = '1'
				s[i] = '1'
				s[i+1] = '0'
			}
		}
	}
	return string(s)
}

func main261() {
	fmt.Println(maximumBinaryStringV1("000110"))
}
