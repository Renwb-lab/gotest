package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/longest-substring-of-one-repeating-character/description/
// 2213. 由单个字符重复的最长子字符串

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	type BeginEnd struct {
		begin, end int
	}
	type Item struct {
		beginEnds []*BeginEnd
		times     int
	}
	helper := make(map[string]*Item)
	pre_c := string(s[0])
	helper[pre_c] = &Item{times: 1, beginEnds: []*BeginEnd{{0, 0}}}
	for i := 1; i < len(s); i += 1 {
		c := string(s[i])
		if c == pre_c {
			item := helper[pre_c]
			item.times += 1
			item.beginEnds[len(item.beginEnds)-1].end = i
		} else {
			if _, ok := helper[c]; !ok {
				helper[c] = &Item{times: 0, beginEnds: []*BeginEnd{}}
			}
			item := helper[c]
			item.times += 1
			item.beginEnds = append(item.beginEnds, &BeginEnd{i, i})
		}
		pre_c = c
	}
	res := make([]int, len(queryCharacters))
	for i, v := range queryIndices {
		// s[v] = queryCharacters[i]
		if i == len(queryCharacters)-1 {
			fmt.Println("test")
		}
		sIdx := v
		oldC, newC := string(s[sIdx]), string(queryCharacters[i])
		{
			// remove
			item := helper[oldC]
			item.times -= 1
			for j, beginEnd := range item.beginEnds {
				if beginEnd.begin <= sIdx && sIdx <= beginEnd.end {
					if beginEnd.begin == sIdx {
						beginEnd.begin += 1
					}
					if beginEnd.end == sIdx {
						beginEnd.end -= 1
					}
					if beginEnd.begin > sIdx && sIdx < beginEnd.end {
						beginEnd.end = sIdx - 1
						item.beginEnds = append(item.beginEnds[:j+1], &BeginEnd{sIdx + 1, beginEnd.end})
						item.beginEnds = append(item.beginEnds, item.beginEnds[j+1:]...)
					}
					if beginEnd.end < beginEnd.begin {
						item.beginEnds = append(item.beginEnds[:j], item.beginEnds[j+1:]...)
					}
					break
				}
			}
			sort.Slice(item.beginEnds, func(i, j int) bool {
				return item.beginEnds[i].begin < item.beginEnds[j].begin ||
					(item.beginEnds[i].begin == item.beginEnds[j].begin && item.beginEnds[i].end == item.beginEnds[j].end)
			})
		}
		{
			item, ok := helper[newC]
			if !ok {
				helper[newC] = &Item{times: 1, beginEnds: []*BeginEnd{{sIdx, sIdx}}}
			} else {
				item.times += 1
				flag := false
				for _, beginEnd := range item.beginEnds {
					if beginEnd.end == sIdx-1 {
						beginEnd.end = sIdx
						flag = true
					}
					if beginEnd.begin == sIdx+1 {
						beginEnd.begin = sIdx
						flag = true
					}
				}
				if flag {
					pre := item.beginEnds[0]
					for j := 1; j < len(item.beginEnds); j += 1 {
						if pre.end == item.beginEnds[j].begin {
							pre.end = item.beginEnds[j].end
							item.beginEnds = append(item.beginEnds[:j], item.beginEnds[j+1:]...)
							break
						} else {
							pre = item.beginEnds[j]
						}
					}
				} else {
					item.beginEnds = append(item.beginEnds, &BeginEnd{sIdx, sIdx})
				}
				sort.Slice(item.beginEnds, func(i, j int) bool {
					return item.beginEnds[i].begin < item.beginEnds[j].begin ||
						(item.beginEnds[i].begin == item.beginEnds[j].begin && item.beginEnds[i].end == item.beginEnds[j].end)
				})
			}
		}

		one := 0
		for _, beginEnds := range helper {
			for _, item := range beginEnds.beginEnds {
				if item.end-item.begin+1 > one {
					one = item.end - item.begin + 1
				}
			}
		}
		res[i] = one
	}
	return res
}

func main124() {
	fmt.Println(longestRepeating("geuqjmt", "bgemoegklm", []int{3, 4, 2, 6, 5, 6, 5, 4, 3, 2}))
}
