package main

import (
	"fmt"
	"math"
)

func minimumDifference(nums []int, k int) int {
	abs := func(x int) int {
		if x < k {
			return k - x
		}
		return x - k
	}
	l := len(nums)
	diff := math.MaxInt
	for i := l - 1; i >= 0; i -= 1 {
		sum := nums[i]
		a := abs(sum)
		if a < diff {
			diff = a
			if diff == 0 {
				return 0
			}
		}
		for j := i - 1; j >= 0; j -= 1 {
			sum = sum & nums[j]
			a = abs(sum)
			if a < diff {
				diff = a
				if diff == 0 {
					return 0
				}
			}
		}
	}
	return diff
}

func minimumDifference1(nums []int, k int) int {
	abs := func(x int) int {
		if x < k {
			return k - x
		}
		return x - k
	}
	l := len(nums)
	diff := math.MaxInt
	for i := 0; i < l; i += 1 {
		if abs(nums[i]) < diff {
			diff = abs(nums[i])
		}
		for j := i - 1; j >= 0; j -= 1 {
			if (nums[j] & nums[i]) == nums[j] {
				break
			}
			nums[j] = nums[j] & nums[i]
			if abs(nums[j]) < diff {
				diff = abs(nums[j])
			}
		}
	}
	return diff
}

func main267() {
	fmt.Println(minimumDifference([]int{1, 2, 4, 5}, 3))
}
