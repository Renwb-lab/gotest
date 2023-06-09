package main

import "fmt"

// https://leetcode.cn/problems/maximum-gap/?company_slug=didi
// 164. 最大间距

func maximumGap(nums []int) (ans int) {
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v > res {
				res = v
			}
		}
		return res
	}

	n := len(nums)
	if n < 2 {
		return
	}

	buf := make([]int, n)
	maxVal := max(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		// 统计每个桶的数量
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		// 进行累加，为下一步的赋值做准备
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		// 重新复制，不清楚为什么要倒序
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}

	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return
}

func maximumGap2(nums []int) (ans int) {
	max := func(args ...int) int {
		r := args[0]
		for _, i := range args[1:] {
			if i > r {
				r = i
			}
		}
		return r
	}
	maxNumber := max(nums...)
	length := len(nums)
	buff := make([]int, length)
	for step := 1; step < maxNumber; step *= 10 {
		bucket := make([]int, 10)
		for _, n := range nums {
			idx := n / step % 10
			bucket[idx] += 1
		}
		for i := 1; i < 10; i += 1 {
			bucket[i] += bucket[i-1]
		}
		for i := length - 1; i >= 0; i -= 1 {
			idx := nums[i] / step % 10
			buff[bucket[idx]-1] = nums[i]
			bucket[idx] -= 1
		}
		copy(nums, buff)
	}
	for i := 1; i < length; i += 1 {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return

}

func main() {
	fmt.Println(maximumGap([]int{
		15252, 16764, 27963, 7817, 26155, 20757, 3478, 22602, 20404, 6739,
		16790, 10588, 16521, 6644, 20880, 15632, 27078, 25463, 20124, 15728,
		30042, 16604, 17223, 4388, 23646, 32683, 23688, 12439, 30630, 3895,
		7926, 22101, 32406, 21540, 31799, 3768, 26679, 21799, 23740,
	}))
	fmt.Println(maximumGap2([]int{
		15252, 16764, 27963, 7817, 26155, 20757, 3478, 22602, 20404, 6739,
		16790, 10588, 16521, 6644, 20880, 15632, 27078, 25463, 20124, 15728,
		30042, 16604, 17223, 4388, 23646, 32683, 23688, 12439, 30630, 3895,
		7926, 22101, 32406, 21540, 31799, 3768, 26679, 21799, 23740,
	}))
}
