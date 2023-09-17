package main

import "fmt"

// 6988. 统计距离为 k 的点对

// 给你一个 二维 整数数组 coordinates 和一个整数 k ，其中 coordinates[i] = [xi, yi] 是第 i 个点在二维平面里的坐标。

// 我们定义两个点 (x1, y1) 和 (x2, y2) 的 距离 为 (x1 XOR x2) + (y1 XOR y2) ，XOR 指的是按位异或运算。

// 请你返回满足 i < j 且点 i 和点 j之间距离为 k 的点对数目。

func countPairs(coordinates [][]int, k int) int {
	r := 0
	l := len(coordinates)
	toString := func(v []int) string {
		return fmt.Sprintf("%d_%d", v[0], v[1])
	}

	m := make(map[string][]int)
	for i := 0; i < l; i += 1 {
		n := 0
		key := toString(coordinates[i])

		if lis, ok := m[key]; ok {
			for j := len(lis) - 1; j >= 0; j -= 1 {
				if lis[j] > i {
					r += 1
				} else {
					break
				}
			}
			continue
		}

		m[key] = make([]int, 0)
		for j := i + 1; j < l; j += 1 {
			t := (coordinates[i][0] ^ coordinates[j][0]) + (coordinates[i][1] ^ coordinates[j][1])
			if t == k {
				n += 1
				m[key] = append(m[key], j)
			}
		}
		fmt.Printf("%+#v", m[key])
		r += n
	}
	return r
}

func main133() {

	coordinates := [][]int{
		{1, 3},
		{1, 3},
		{1, 3},
		{1, 3},
		{1, 3},
	}
	fmt.Print(countPairs(coordinates, 0))
}
