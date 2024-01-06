package main

import "fmt"

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	helper := make([]uint, 1<<21)
	for i := range helper {
		helper[i] = 0
	}

	var dfs func(state, total int) bool
	dfs = func(state, total int) bool {
		if helper[state] == 1 {
			return true
		}
		if helper[state] == 2 {
			return false
		}
		for i := 1; i <= maxChoosableInteger; i += 1 {
			if ((1 << i) & state) > 0 {
				continue
			}
			if i+total >= desiredTotal || !dfs(state|(1<<i), total+i) {
				helper[state] = 1
				return true
			}
		}
		helper[state] = 2
		return false
	}

	return dfs(0, 0)
}

func main() {
	fmt.Println(canIWin(10, 40))
}
