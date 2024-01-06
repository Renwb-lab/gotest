package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Solution struct {
	radius  float64
	xCenter float64
	yCenter float64
	r       float64
	r2      float64
}

func Constructor(radius float64, xCenter float64, yCenter float64) Solution {
	return Solution{
		radius:  radius,
		xCenter: xCenter,
		yCenter: yCenter,
	}
}

func (this *Solution) RandPoint() []float64 {
	// r := math.Sqrt(rand.Float64())
	// this.r = math.Sqrt(rand.Float64())
	this.r = rand.Float64()
	this.r2 = rand.Float64() * 2 * math.Pi
	sin, cos := math.Sincos(this.r2)
	return []float64{
		this.xCenter + this.r*cos*this.radius,
		this.yCenter + this.r*sin*this.radius,
	}
}

func main() {
	s := Constructor(0.01, -73839.1, -3289891.3)
	i := 0
	for {
		if i%100000 == 0 {
			fmt.Println(i)
		}
		res := s.RandPoint()
		if (res[0]-s.xCenter)*(res[0]-s.xCenter)+(res[1]-s.yCenter)*(res[1]-s.yCenter) > s.radius*s.radius {
			a, b := res[0]-s.xCenter, res[1]-s.yCenter
			fmt.Println(a, b, s.radius)
			fmt.Println(s.r, s.r2)
			sin, cos := math.Sincos(s.r2)
			fmt.Println(sin, cos)
			break
		}
		i += 1
	}

}
