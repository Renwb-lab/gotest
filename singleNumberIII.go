package main

// https://leetcode.cn/problems/single-number-iii
func singleNumberIII(nums []int) []int {
	OrSum := func() int {
		x := nums[0]
		for _, v := range nums[1:] {
			x = x ^ v
		}
		return x
	}

	firstOne := func(xor int) int {
		for i := 0; i < 32; i += 1 {
			if (xor & (1 << i)) > 0 {
				return i
			}
		}
		return 0
	}

	getNumber := func(bit int) (int, int) {
		x, y := 0, 0
		for _, v := range nums {
			if v&(1<<bit) > 0 {
				x = x ^ v
			} else {
				y = y ^ v
			}
		}
		if x > y {
			return y, x
		}
		return x, y
	}

	xor := OrSum()
	leftBit := firstOne(xor)
	one, two := getNumber(leftBit)
	return []int{one, two}
}
