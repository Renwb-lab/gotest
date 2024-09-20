package main

// （200分）有一个字符串 ，如 “AABBA”，（只有 ‘A’ 和 ‘B’  组成 ，其中规定 ‘B’  > ‘A’），
//  你可以改变其中任意一个字符  （A -> B 或 B -> A），使其变成一个递增串 ，求把该字符串变成递增字符串的最少次数
// 	示例 1：
// 	输入 ： “AABBA”
// 	输出 ：1
// 	解释 ： 只需要把最后一个 A 改成 B ，变成 “AABBB” ，此时这个字符串递增 ，只需要变1次

func ABReverse(inStr string) int {
	// 转化之后，最终的形式为A...A[n个]B...B[m]个， 其中 n + m 等于字符串的长度
	// 也就是说只要查找一个出来一个位置，该位置前左边的字符都A, 右边的字符都是B, 计算更改的数量就可以了。
	// 遍历每一个位置，计算改变量，求最小值即可。

	// 优化，两个特例
	// 从左边开始找连续的A，然后计算该位置pos1后全部为A的改变量num1,
	// 从右边开始找连续的B，然后计算该位置pos2前全部为B的改变量num2,
	// 返回num1, num2的最小值

	pos1, num1, l := 0, 0, len(inStr)
	for ; pos1 < l && inStr[pos1] == 'A'; pos1 += 1 {
	}
	for ; pos1 < l; pos1 += 1 {
		if inStr[pos1] == 'A' {
			num1 += 1
		}
	}
	if num1 == 0 {
		return num1
	}
	pos2, num2 := l-1, 0
	for ; pos2 >= 0 && inStr[pos2] == 'B'; pos2 -= 1 {

	}
	for ; pos2 >= 0; pos2 -= 1 {
		if inStr[pos2] == 'B' {
			num2 += 1
		}
	}
	if num1 < num2 {
		return num1
	}
	return num2
}
