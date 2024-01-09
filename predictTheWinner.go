package main

import "fmt"

// https://leetcode.cn/problems/predict-the-winner/description/
func predictTheWinner(nums []int) bool {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// dp
	l := len(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	var dp func(x, y int) int
	dp = func(x, y int) int {
		if x >= l || y < 0 {
			return 0
		}
		if x > y {
			return 0
		}
		if x == y {
			return nums[x]
		}
		t1 := nums[x] + min(dp(x+2, y), dp(x+1, y-1))
		t2 := nums[y] + min(dp(x, y-2), dp(x+1, y-1))
		return max(t1, t2)
	}
	s := dp(0, l-1)
	return 2*s >= sum
}

func predictTheWinner2(nums []int) bool {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// dp
	l := len(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	mp := make([][]int, l)
	for i := range mp {
		mp[i] = make([]int, l)
		for j := range mp[i] {
			mp[i][j] = -1
		}
	}
	var dp func(x, y int) (res int)
	dp = func(x, y int) (res int) {
		if x >= l || y < 0 {
			res = 0
			return
		}
		if mp[x][y] != -1 {
			res = mp[x][y]
		} else if x > y {
			res = 0
		} else if x == y {
			res = nums[x]
		} else {
			t1 := nums[x] + min(dp(x+2, y), dp(x+1, y-1))
			t2 := nums[y] + min(dp(x, y-2), dp(x+1, y-1))
			res = max(t1, t2)
		}
		mp[x][y] = res
		return
	}
	s := dp(0, l-1)
	return 2*s >= sum
}

func predictTheWinner3(nums []int) bool {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// dp
	l := len(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	mp := make([][]int, l)
	for i := range mp {
		mp[i] = make([]int, l)
		for j := range mp[i] {
			mp[i][j] = -1
		}
	}

	for i := 0; i < l; i += 1 {
		mp[i][i] = nums[i]
	}
	// 1, 5, 3,
	// 1  5  0
	// 0  5  5
	// 0  0  3

	get := func(x, y int) int {
		if x > y {
			return 0
		}
		return mp[x][y]
	}
	for x := l - 2; x >= 0; x -= 1 {
		for y := x + 1; y < l; y += 1 {
			t1 := nums[x] + min(get(x+2, y), get(x+1, y-1))
			t2 := nums[y] + min(get(x, y-2), get(x+1, y-1))
			mp[x][y] = max(t1, t2)
		}
	}

	return 2*mp[0][l-1] >= sum
}

func main() {
	fmt.Println(predictTheWinner3([]int{1, 5, 3, 7}))
}
