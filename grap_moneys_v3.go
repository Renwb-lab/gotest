package main

import (
	"fmt"
	"math"
	"math/rand"
)

type RedPackegeV2 struct {
	remainMoney int32
	remainSize  int32
}

func NewRedPackegeV2(remainMoney int32, remainSize int32) *RedPackegeV2 {
	return &RedPackegeV2{
		remainMoney: remainMoney * 100,
		remainSize:  remainSize,
	}
}

func (rp *RedPackegeV2) GetRandomMoney() float64 {
	if rp.remainSize == 1 {
		money := rp.remainMoney
		rp.remainMoney -= money
		rp.remainSize -= 1
		return float64(money) / 100.0
	}
	avg := int32(math.Round(float64(rp.remainMoney) / float64(rp.remainSize)))
	money := rand.Int31n(avg*2) + 1 // rand.Float64() * max
	rp.remainMoney -= money
	rp.remainSize -= 1
	return float64(money) / 100.0
}

func main() {
	rp := NewRedPackegeV2(100, 5)
	sumMoney := float64(0.0)
	for i := 0; i < 5; i += 1 {
		money := rp.GetRandomMoney()
		fmt.Println(money)
		sumMoney += money
	}
	fmt.Println(sumMoney)
}
