package main

import "fmt"

func getRow(rowIndex int) []int {
	// A(m, n) = n * (n-1) * (n-2) * .... * (n - m + 1)
	// C(m, n) = n * (n-1) * (n-2) * .... * (n - m + 1) / m!
	// n! = n * (n-1) * (n-2) * .... * 1
	// (n-m)!= (n-m) * (n-m-1) * .... * 1
	// C(m, n) = n! / (m! * (n-m)!)
	// C(m-1, n) = n! / ((m-1)! * (n-m+1)!)

	// C(m, n) = C(m-1, n) * n! / (m! * (n-m)!) / (n! / ((m-1)! * (n-m+1)!))
	//         = C(m-1, n) * ((m-1)! * (n-m+1)!) / (m! * (n-m)!)
	//         = C(m-1, n) * (m-1)!/m! * (n-m+1)!/(n-m)!
	//         = C(m-1, n) * 1/m * (n-m+1)
	//         = C(m-1, n) * (n-m+1) / m
	res := make([]int, rowIndex+1, rowIndex+1)
	res[0], res[rowIndex] = 1, 1
	for i := 1; i < rowIndex; i += 1 {
		// 1, 4, 6
		res[i] = res[i-1] * (rowIndex - i + 1) / i
	}
	return res
}

func main() {
	fmt.Println(getRow(3))
}
