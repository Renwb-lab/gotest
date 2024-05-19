package main

import (
	"fmt"
)

var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var sf2mask = [31]int{} // sf2mask[i] 为 i 的质因子集合（用二进制表示）

func init() {
	for i := 2; i <= 30; i++ {
		for j, p := range primes {
			if i%p == 0 {
				if i%(p*p) == 0 { // 有平方因子
					sf2mask[i] = -1
					break
				}
				sf2mask[i] |= 1 << j // 把 j 加到集合中
			}
		}
	}
}

func squareFreeSubsets(nums []int) int {
	const mod int = 1e9 + 7
	const m = 1 << len(primes)
	f := [m]int{1} // f[j] 表示恰好组成质数集合 j 的方案数，其中质数集合是空集的方案数为 1
	for _, x := range nums {
		if mask := sf2mask[x]; mask >= 0 { // x 是 SF
			for j := m - 1; j >= mask; j-- {
				if j|mask == j { // mask 是 j 的子集
					f[j] = (f[j] + f[j^mask]) % mod // 不选 mask + 选 mask
				}
			}
		}
	}
	ans := 0
	for _, v := range f {
		ans += v
	}
	return (ans - 1) % mod // -1 去掉空集（nums 的空子集）
}

func main() {
	fmt.Println(squareFreeSubsets([]int{3, 5, 21, 20, 10,
		2, 29, 18, 28, 1, 29, 15, 18, 3, 22, 19, 3, 14,
		22, 15, 18, 13, 12, 26, 12, 26, 17, 10, 6, 19, 21,
		14, 10, 26, 18, 19, 20, 28, 12, 12, 15, 28, 19, 13,
		20, 17}))
}
