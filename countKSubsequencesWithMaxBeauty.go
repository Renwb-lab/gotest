package main

import (
	"fmt"
	"sort"
)

const mod = 1000000007

func countKSubsequencesWithMaxBeauty(s string, k int) int {
	// 1. 统计每个字符出现的次数 [a: 0, b: 1, c: 2, d: 2, e:2]
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a'] += 1
	}
	// 2. 统计出现次数的次数 [0: 1, 1: 1, 2: 3]
	cc := map[int]int{}
	for _, c := range cnt {
		if c > 0 {
			cc[c]++
		}
	}
	// 3. 根据次数进行排序
	type KV struct{ cnt, num int }
	kv := make([]KV, 0, len(cc))
	for k, v := range cc {
		kv = append(kv, KV{k, v})
	}
	sort.Slice(kv, func(i, j int) bool { return kv[i].cnt > kv[j].cnt })
	// 4. 使用数学方法进行求解
	ans := 1
	for _, p := range kv {
		if p.num >= k {
			return ans * pow(p.cnt, k) % mod * comb(p.num, k) % mod
		}
		ans = ans * pow(p.cnt, p.num) % mod
		k -= p.num
	}
	return 0 // k 太大，无法选 k 个不一样的字符
}

func pow(x, n int) int {
	ans := 1
	// pow(5, 3)
	for ; n > 0; n = n >> 1 {
		if n&1 > 0 {
			ans = ans * x % mod
		}
		x = x * x % mod
	}
	return ans
}

func comb(n, k int) int {
	// comb(n, k) =
	// n * (n - 1) * (n - 2)  * ... * (n-k+1) / 1 * 2 * ... * k
	res := n
	for i := 2; i <= k; i += 1 {
		res = res * (n - i + 1) / i
	}
	return res % mod
}

func main() {
	fmt.Println(countKSubsequencesWithMaxBeauty("bcca", 2))
}
