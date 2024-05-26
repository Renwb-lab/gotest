package main

import "fmt"

func longestAwesome(s string) int {
	status, ans := 0, 1
	prefix := map[int]int{status: -1}
	// status 存放0-9数字出现的次数，奇数设置为1，偶数设置为0
	// prefix 保存出现status时的最左边下标
	// 如果前后出现了相同的status， 那么表明中间出现的字节数量是偶数，满足条件
	// 如果前后出现了只有一位不同的status， 那么表明中间出现的字节数量是奇数，满足条件
	for i, c := range s {
		n := byte(c) - byte('0')
		status = status ^ (1 << n)

		// 只有一位不同的情况
		for j := 0; j < 10; j += 1 {
			ns := status ^ (1 << j)
			if p, ok := prefix[ns]; ok {
				if i-p+1 > ans {
					ans = i - p
				}
			}
		}

		// 出现了相同的状态
		if p, ok := prefix[status]; ok {
			if i-p+1 > ans {
				ans = i - p
			}
		} else {
			prefix[status] = i
		}
	}
	return ans
}

func main() {
	fmt.Println(longestAwesome("24142"))
}
