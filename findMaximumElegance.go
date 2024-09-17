package main

import (
	"fmt"
	"sort"
)

// 2813. 子序列最大优雅度
// 这里只是计算子序列的最大优雅度，只需要计算结果就可以。
// 因此子序列可以理解为只要选择k个就可以了。
// 如果需要输出的k个，就需要在排序中添加下标
func findMaximumElegance(items [][]int, k int) int64 {
	// 按照最大价值进行排序
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})
	categorys := make(map[int]struct{})
	duplicate := []int{} // 重复类别的利润
	ans, tmp := 0, 0
	// 先选择前K个
	for i, p := range items {
		fit, c := p[0], p[1]
		if i < k {
			tmp += fit
			// 标记类别是否出现过
			if _, ok := categorys[c]; !ok {
				categorys[c] = struct{}{} // 没有出现过的话，则就类别添加进去
			} else {
				duplicate = append(duplicate, fit) // 记录重复出现类别的价值
			}
		} else if _, ok := categorys[c]; ok {
			// 如果类别出现过，那么category就不会变大，而fit还会变小，因此直接忽略当前选项
		} else if _, ok := categorys[c]; !ok && len(duplicate) > 0 {
			// 如果类别没有出现过，那么category就会变大，而fit选择最小进行移除
			categorys[c] = struct{}{}
			tmp += fit - duplicate[len(duplicate)-1]
			duplicate = duplicate[:len(duplicate)-1]
		}
		if tmp+len(categorys)*len(categorys) > ans {
			ans = tmp + len(categorys)*len(categorys)
		}
	}
	return int64(ans)
}

func main071042() {
	fmt.Println(findMaximumElegance([][]int{
		{3, 2},
		{5, 1},
		{10, 1},
	}, 2))
}
