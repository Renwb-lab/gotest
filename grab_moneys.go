package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func grab_moneys(money float32, people int) []float32 {
	if money*100 < float32(people) {
		return nil
	}
	if people < 1 {
		return nil
	}
	if people == 1 {
		return []float32{money}
	}
	upBound := int32(money * 100)
	randomList := make([]int32, people+1, people+1)
	randomList[0], randomList[people] = 0, upBound

	r := rand.New(rand.NewSource(time.Now().Local().Unix()))
	getRandom := func() int32 {
		for {
			n := r.Int31n(upBound-1) + 1
			flag := true
			for i := range randomList {
				if randomList[i] == n {
					flag = false
				}
			}
			if flag {
				return n
			}
		}
	}
	// 100 5 -> [1--99]
	// [1,2,3,4, 100]

	for i := 1; i < people; i += 1 {
		randomList[i] = getRandom()
	}
	sort.Slice(randomList, func(i, j int) bool { return randomList[i] < randomList[j] })

	res := make([]float32, people, people)
	for i := 1; i <= people; i += 1 {
		res[i-1] = float32(randomList[i]-randomList[i-1]) / 100.0
	}
	return res
}

func main071050() {
	fmt.Println(grab_moneys(1, 99))
}
