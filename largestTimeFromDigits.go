package main

import "fmt"

func largestTimeFromDigits(arr []int) string {
	res, maxTime := "", -1

	compute := func() int {
		hour := arr[0]*10 + arr[1]
		minute := arr[2]*10 + arr[3]
		if hour < 24 && minute < 60 {
			return hour*60 + minute
		}
		return -1
	}
	getRes := func() string {
		return fmt.Sprintf("%d%d:%d%d", arr[0], arr[1], arr[2], arr[3])
	}

	l := len(arr)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == l {
			thisTime := compute()
			if thisTime > maxTime {
				res = getRes()
				maxTime = thisTime
			}
		}
		for i := idx; i < l; i += 1 {
			arr[i], arr[idx] = arr[idx], arr[i]
			dfs(idx + 1) // 特别重要
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	dfs(0)
	return res
}

func main246() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
}
