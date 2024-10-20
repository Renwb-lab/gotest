package main

import "math"

func twoEggDrop(n int) int {
	// f[i] 表示 i 层楼的建筑需要的最小的操作次数。
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	// 如果鸡蛋碎了，说明答案的范围是 [0,k−1]
	// 如果鸡蛋没碎，说明答案的范围是 [k,i]，并且我们还剩下两枚鸡蛋
	for i := 1; i <= n; i += 1 {
		for k := 1; k <= i; k += 1 {
			// 破了：k-1
			// 没破：f[i-k]
			f[i] = min(f[i], max(k-1, f[i-k])+1)
		}
	}
	return f[n]
}

// f[x] = y ==> x层楼最小操作次数
// 		= k - 1   			// 在k楼破了
// 		= f[x-k]			// 在k楼没破
// 破的楼层不确定，需要取最大值
