package main

import "fmt"

func printBin(num float64) string {
	ret := []byte{'0', '.'}
	n, t := 0.5, 0
	for ; num != 0 && t <= 32; t += 1 {
		if num >= n {
			ret = append(ret, '1')
			num -= n
		} else {
			ret = append(ret, '0')
		}
		n = n / 2
	}
	if num == 0 {
		return string(ret)
	}
	return "ERROR"
}

func main14() {
	fmt.Println(printBin(0.625))
	fmt.Println(printBin(0.1))
}
