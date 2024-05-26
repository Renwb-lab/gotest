package main

import "fmt"

func mostCompetitive(nums []int, k int) []int {
	better := func(a []int, b []int, k int) bool {
		for i := 0; i < k; i += 1 {
			if a[i] < b[i] {
				return true
			}
		}
		return false
	}

	ans := []int{}
	l := len(nums)
	var dp func(idx int, helper []int)
	dp = func(idx int, helper []int) {
		if idx == l {
			if len(helper) == k && (len(ans) == 0 || better(helper, ans, k)) {
				ans = append([]int{}, helper...)
			}
			return
		}
		if len(helper) > k {
			return
		}
		if len(helper)+l-idx < k {
			return
		}

		if len(ans) > 0 {
			for i, v := range helper {
				if ans[i] < v {
					return
				} else if ans[i] > v {
					break
				}
			}
		}

		if len(helper) < k {
			helper = append(helper, nums[idx])
			dp(idx+1, helper)
			helper = helper[:len(helper)-1]
		}

		dp(idx+1, helper)
	}
	dp(0, []int{})
	return ans
}

func mostCompetitive2(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	for i, x := range nums {
		for len(res) > 0 &&
			len(nums)-i+len(res) > k && // 注意这里是 大于k, 不能是等于，因为等于的话，后面还需要出一个
			res[len(res)-1] > x {
			res = res[:len(res)-1]
		}
		res = append(res, x)
	}
	return res[:k]
}

func main() {
	fmt.Println(mostCompetitive([]int{
		2, 10, 3, 5, 9, 4, 2, 0, 6, 7, 8, 0, 6,
		5, 8, 1, 6, 1, 5, 5, 2, 10, 9, 5, 7, 7,
		3, 2, 1, 4, 0, 7, 0, 3, 10, 10, 5, 10,
		4, 7, 0, 2, 10, 9, 0, 2, 6, 10, 6, 9,
		2, 1, 9, 8, 7, 2, 0, 7, 3, 6, 2, 1,
		8, 0, 0, 0, 10, 4, 3, 5, 0, 8, 1, 8,
		5, 1, 6, 0, 4, 4, 10, 2, 0, 5, 1, 1,
		3, 3, 5, 2, 6, 5, 6, 0, 3, 8, 0, 1,
		7, 0, 0, 9, 6, 10, 5, 9, 8, 9, 8, 7,
		8, 10, 6, 3, 8, 0, 5, 7, 4, 3, 5, 7,
		7, 0, 3, 10, 1, 3, 10, 2, 10, 3, 2, 6,
		3, 10, 8, 10, 6, 0, 7, 6, 2, 10, 4, 0,
		7, 4, 8, 8, 1, 7, 1, 4, 9, 7, 7, 8, 9, 8,
		7, 2, 4, 9, 8, 8, 0, 8, 2, 10, 7, 3, 10,
		8, 5, 1, 1, 3, 0, 5, 1, 7, 1, 7, 9, 2, 6,
		9, 6, 10, 6, 1, 7, 8, 3, 6, 9, 3, 5, 9,
		0, 9, 3, 5, 8, 4, 6, 8, 10, 8, 0, 9, 3,
		7, 10, 4, 4, 2, 3, 7, 2, 10, 3, 5, 4, 9, 9,
		2, 1, 2, 10, 4, 4, 4, 3, 5, 9, 7, 2, 0,
		3, 6, 6, 7, 3, 9, 4, 6, 9, 7, 1, 3, 2,
		3, 6, 6, 1, 7, 10, 0, 4, 10, 3, 5, 0,
		10, 3, 10, 3, 0, 0, 1, 6, 6, 5, 9, 10,
		5, 5, 9, 0, 5, 4, 1, 10, 2, 3, 1, 7,
		9, 10, 10, 4, 3, 5, 9, 5, 4, 4, 8, 0,
		1, 8, 1, 4, 6, 5, 6, 0, 6, 8, 6, 5, 6,
		5, 7, 9, 5, 8, 8, 4, 2, 0, 0, 2, 9,
		4, 9, 2, 6, 5, 2, 2, 8, 5, 4, 10,
		8, 7, 7, 3, 4, 2, 0, 4, 3, 8, 6, 1,
		7, 10, 10, 7, 4, 0, 6, 6, 0, 5, 6,
		10, 3, 8, 3, 2, 4, 10, 4, 3, 0, 4,
		10, 7, 6, 0, 4, 7, 0, 5, 2, 5, 2, 10, 9,
		1, 10, 9, 6, 6, 5, 9, 10, 1, 3, 5, 2, 0,
		6, 8, 5, 6, 3, 4, 8, 4, 0, 7, 0, 7, 9, 9,
		1, 4, 6, 4, 5, 7, 3, 0, 4, 4, 9, 10, 5, 10,
		3, 9, 6, 6, 2, 9, 4, 0, 4, 3, 3, 1, 7, 2,
		1, 0, 2, 6, 7, 1, 1, 0, 3, 9, 8, 9, 4, 6,
		3, 10, 7, 3, 1, 5, 2, 0, 3, 9, 5, 3, 3,
		3, 1, 7, 5, 8, 10, 10, 8, 0, 2, 3, 3, 2,
		9, 3, 1, 3, 9, 0, 1, 8, 2, 1, 6, 0, 6,
		3, 1, 3, 1, 10, 5, 6, 0, 4, 7, 10}, 24))
}
