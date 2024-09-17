package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	// 1. 0 + n + x
	// 2. 1 * n + x
	// 3. 2 * n + x
	// 4. 3 * n + x
	// ...
	// n. y * n + x

	// (1 + 2 + ... + y) * n + (y + 1) * n
	// (1 + y) * y / 2 * n + (y + 1) * n

	n := (1 + num_people) * num_people / 2
	t := 0
	for y := 1; ; y += 1 {
		x := (1+y)*y/2*n + (y+1)*n
		if x > candies {
			t = y - 1
			break
		}
	}

	ans := make([]int, num_people)
	extra := candies - (1+t)*t/2*n - (t+1)*n
	for i := 0; i < num_people; i += 1 {
		ans[i] = (num_people+1)*t + i + 1
		if extra > 0 {
			v := (t + 1) * (num_people + 1)
			if extra > v {
				ans[i] += v
				extra -= v
			} else {
				ans[i] += extra
				extra = 0
			}
		}
	}

	return ans
}

func main071035() {
	fmt.Println(distributeCandies(10, 3))
}
