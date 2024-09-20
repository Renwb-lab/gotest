package main

import (
	"fmt"
	"sort"
)

func findOriginalArrayV1(changed []int) []int {
	l := len(changed)
	if l&0x1 == 1 {
		return []int{}
	}
	check := func(min, max int) bool {
		t1, t2 := min, max
		for t1 > 0 && t2 > 0 {
			t1 = t1 & (t1 - 1)
			t2 = t2 & (t2 - 1)
		}
		if t1 != 0 || t2 != 0 {
			return false
		}
		for max > min {
			max = max >> 1
		}
		return max == min
	}
	sort.Slice(changed, func(i, j int) bool { return changed[i] < changed[j] })
	m := map[int]map[int]int{}
	mh := map[int][]int{}
	for _, n := range changed {
		flag := false
		idx := 0
		for k := range m {
			if check(k, n) {
				flag, idx = true, k
				break
			}
		}
		if flag {
			if _, ok := m[idx][n/2]; ok {
				m[idx][n] += 1
				mh[idx] = append(mh[idx], n)
			} else if _, ok := m[idx][n]; ok {
				m[idx][n] += 1
				mh[idx] = append(mh[idx], n)
			}
		} else {
			m[n] = map[int]int{}
			m[n][n] += 1
			mh[n] = append(mh[n], n)
		}
	}
	res := []int{}
	for k, l := range mh {
		if len(l)&1 == 1 {
			return []int{}
		}
		sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
		t := []int{}
		for i := 0; i < len(l)-1; i += 1 {
			if len(t) >= len(l)/2 {
				break
			}
			flag := false
			if m[k][l[i]<<1] > 0 {
				m[k][l[i]<<1] -= 1
				flag = true
			}
			if m[k][l[i]>>1] > 0 {
				m[k][l[i]>>1] -= 1
				flag = true
			}
			if !flag {
				return []int{}
			}
			if (l[i] << 1) == l[i+1] {
				t = append(t, l[i])
				i += 1
			} else {
				t = append(t, l[i])
			}
		}
		res = append(res, t...)
	}
	return res
}

func findOriginalArray(changed []int) []int {
	sort.Slice(changed, func(i, j int) bool { return changed[i] < changed[j] })
	m := map[int]int{}
	for _, n := range changed {
		m[n] += 1
	}
	res := []int{}
	for _, n := range changed {
		if m[n] == 0 {
			continue
		}
		m[n] -= 1
		if m[2*n] == 0 {
			return []int{}
		}
		m[2*n] -= 1
		res = append(res, n)
	}
	return res
}

func main071043() {
	fmt.Println(findOriginalArray([]int{52184, 36160, 76173, 4238, 88165, 71309, 53758, 63977, 83042, 68805, 12366, 27457, 65992, 88573, 18449, 98384}))
}
