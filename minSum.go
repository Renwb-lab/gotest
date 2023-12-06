package main

import "fmt"

func minSum(nums1 []int, nums2 []int) int64 {
	f := func(nums []int) (s int64, z int64) {
		for _, n := range nums {
			if n == 0 {
				z += 1
			}
			s += int64(n)
		}
		return
	}
	max := func(n1, n2 int64) int64 {
		if n1 < n2 {
			return n2
		}
		return n1
	}
	s1, z1 := f(nums1)
	s2, z2 := f(nums2)
	if s1 < s2 {
		s1, s2 = s2, s1
		z1, z2 = z2, z1
	}
	if s1 == s2 {
		if z1 == 0 && z2 == 0 {
			return s1
		} else if z1 == 0 && z2 != 0 {
			return -1
		} else if z1 != 0 && z2 == 0 {
			return -1
		} else {
			return max(s1+z1, s2+z2)
		}
	} else {
		if z1 == 0 && z2 == 0 {
			return -1
		} else if z1 == 0 && z2 != 0 {
			if s1 >= s2+z2 {
				return s1
			} else {
				return -1
			}
		} else if z1 != 0 && z2 == 0 {
			return -1
		} else {
			return max(s1+z1, s2+z2)
		}
	}
}

func main157() {
	fmt.Println(minSum(
		[]int{20, 0, 18, 11, 0, 0, 0, 0, 0, 0, 17, 28, 0, 11, 10, 0, 0, 15, 29},
		[]int{16, 9, 25, 16, 1, 9, 20, 28, 8, 0, 1, 0, 1, 27},
	))
}
