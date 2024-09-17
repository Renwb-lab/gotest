package main

import (
	"fmt"
	"math"
)

func maxScore(nums []int, x int) int64 {
	l := len(nums)
	helper := make([]int, l, l) // 表示以当前值作为最后一个添加元素时的最大值
	helper[0] = nums[0]
	ans := helper[0]
	for i := 1; i < l; i += 1 {
		helper[i] = math.MinInt32
		for j := 0; j < i; j += 1 {
			t := 0
			if (nums[j] % 2) == (nums[i] % 2) {
				t = helper[j] + nums[i]
			} else {
				t = helper[j] + nums[i] - x
			}
			if t > helper[i] {
				helper[i] = t
			}
		}
		if helper[i] > ans {
			ans = helper[i]
		}
	}
	return int64(ans)
}

func maxScore2(nums []int, x int) int64 {
	// [0:i]个元素组成的数组的返回值
	// a 为偶数作为结果的最大值， b为奇数作为结果的最大值
	a, b := math.MinInt32, math.MinInt32
	if nums[0]%2 == 0 {
		a = nums[0]
	} else {
		b = nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < len(nums); i += 1 {
		if nums[i]%2 == 0 {
			t1 := a + nums[i]
			t2 := b + nums[i] - x
			a = max(t1, t2)
		} else {
			t1 := a + nums[i] - x
			t2 := b + nums[i]
			b = max(t1, t2)
		}
	}
	return int64(max(a, b))
}

func main264() {
	fmt.Println(maxScore2(
		[]int{
			8, 50, 65, 85, 8, 73, 55, 50, 29, 95, 5, 68, 52, 79,
		},
		74))
}
