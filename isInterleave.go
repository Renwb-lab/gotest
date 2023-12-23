package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	var dp func(idx1 int, idx2 int, idx3 int) bool
	dp = func(idx1 int, idx2 int, idx3 int) bool {
		if idx1 == n1 && idx2 == n2 && idx3 == n3 {
			return true
		}
		if s3[idx3] == s1[idx1] {
			if dp(idx1, idx2+1, idx3+1) {
				return true
			}
		}
		if s3[idx3] == s2[idx2] {
			if dp(idx1+1, idx2, idx3+1) {
				return true
			}
		}
		return false
	}
	r := dp(0, 0, 0)
	return r
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	// helper[i][j] 表示s1的前i元素和s2的前j元素能否组成s3的前i+j元素
	helper := make([][]bool, n1+1, n1+1)
	for i := 0; i < n1+1; i += 1 {
		helper[i] = make([]bool, n2+1, n2+1)
	}
	helper[0][0] = true
	// helper[0][1]=true
	// 表示s1的前i元素 “” 和s2的前j元素 s2[0] 能组成s3的前i+j元素 s3[0], 即s2[0] = s3[0]
	for i := 0; i < n1+1; i += 1 {
		for j := 0; j < n2+1; j += 1 {
			p := i + j
			if i > 0 && s1[i-1] == s3[p-1] {
				helper[i][j] = helper[i][j] || helper[i-1][j]
			}
			if j > 0 && s2[j-1] == s3[p-1] {
				helper[i][j] = helper[i][j] || helper[i][j-1]
			}
		}
	}
	return helper[n1][n2]
}

func isInterleave3(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	// helper[i][j] 表示s1的前i元素和s2的前j元素能否组成s3的前i+j元素
	helper := make([]bool, n2+1, n2+1)
	helper[0] = true
	// helper[0][1]=true
	// 表示s1的前i元素 “” 和s2的前j元素 s2[0] 能组成s3的前i+j元素 s3[0], 即s2[0] = s3[0]
	for i := 0; i < n1+1; i += 1 {
		for j := 0; j < n2+1; j += 1 {
			p := i + j
			if i > 0 {
				if s1[i-1] != s3[p-1] {
					helper[j] = false
				}
			}
			if j > 0 {
				if s2[j-1] == s3[p-1] {
					helper[j] = helper[j] || helper[j-1]
				}
			}
		}
	}
	return helper[n2]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave2("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave3("aabcc", "dbbca", "aadbbbaccc"))
}
