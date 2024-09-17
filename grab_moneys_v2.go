package main

import (
	"fmt"
	"math"
	"math/rand"
)

type RedPackege struct {
	remainMoney float64
	remainSize  int
}

func NewRedPackege(remainMoney float64, remainSize int) *RedPackege {
	return &RedPackege{
		remainMoney: remainMoney,
		remainSize:  remainSize,
	}
}

func (rp *RedPackege) GetRandomMoney() float64 {
	if rp.remainSize == 1 {
		money := rp.remainMoney
		rp.remainMoney -= money
		rp.remainSize -= 1
		return float64(math.Round(money*100)) / 100.0
	}
	avg := rp.remainMoney / float64(rp.remainSize)
	min, max := float64(0.01), avg*2
	money := rand.Float64() * max
	if money < min {
		money = min
	}
	money = float64(math.Round(money*100)) / 100.0 // 报错两位小数
	rp.remainMoney -= money
	rp.remainSize -= 1
	return money
}

func main253() {
	for i := 0; i < 100; i += 1 {
		rp := NewRedPackege(100, 5)
		sumMoney := float64(0.0)
		for i := 0; i < 5; i += 1 {
			money := rp.GetRandomMoney()
			// fmt.Println(money)
			sumMoney += money
		}
		fmt.Println(sumMoney)
	}

}
