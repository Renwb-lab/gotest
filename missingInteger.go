package main

import "fmt"

func missingInteger(nums []int) int {
	if len(nums) == 1 {
		return nums[0] + 1
	}

	sum := 0
	ll := 0
	m := map[int]struct{}{
		nums[0]: struct{}{},
	}
	for j := 1; j < len(nums); {
		t := 1
		ts := nums[j-1]
		if nums[j]-nums[j-1] == 1 {
			for j < len(nums) && nums[j]-nums[j-1] == 1 {
				ts += nums[j]
				m[nums[j]] = struct{}{}
				j, t = j+1, t+1
			}
		} else {
			m[nums[j]] = struct{}{}
			j += 1
		}

		if t > ll {
			ll = t
			sum = ts
		}
	}
	for {
		if _, ok := m[sum]; !ok {
			return sum
		}
		sum += 1
	}
}

func main280() {
	fmt.Println(missingInteger([]int{3, 4, 5, 1, 12, 14, 13}))
}
