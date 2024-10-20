package main

func TwoSum(arr []int, target int) []int {
	m := make(map[int]int)
	for i := range arr {
		diff := target - arr[i]
		if preIdx, ok := m[diff]; ok {
			return []int{preIdx, i}
		}
		m[arr[i]] = i
	}
	return []int{-1, -1}
}
