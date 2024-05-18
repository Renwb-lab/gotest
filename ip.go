package main

import (
	"fmt"
	"strconv"
	"strings"
)

func judgeIp(ip string) bool {
	v := strings.Split(ip, ".")
	if len(v) != 4 {
		return false
	}
	for _, s := range v {
		n, e := strconv.Atoi(s)
		if e != nil || n > 255 {
			return false
		}
	}
	return true
}

func isPrivate(ip string) bool {
	v := strings.Split(ip, ".")
	if len(v) != 4 {
		return false
	}
	first, _ := strconv.Atoi(v[0])
	second, _ := strconv.Atoi(v[1])
	return first == 0 ||
		(first == 172 && second >= 16 && second <= 31) ||
		(first == 192 && second == 168)
}

func isMask(ip string) bool {
	v := strings.Split(ip, ".")
	if len(v) != 4 {
		return false
	}
	var b uint32 = 0
	for _, s := range v {
		n, _ := strconv.Atoi(s)
		b = b<<8 + uint32(n)
	}
	if b == 0 {
		return false
	}
	b = ^b + 1
	if b == 1 {
		return false
	}
	return (b & (b - 1)) == 0
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	var A string
	a, b, c, d, e, f, g := 0, 0, 0, 0, 0, 0, 0
	for _, line := range []string{
		"110.156.20.173~255.153.0.0",
		"225.240.129.203~255.110.255.255",
		"183.181.49.4~255.0.0.0",
		"172.177.113.45~255.0.0.0",
		"176.134.46.246~255.0.0.0",
		"153.63.21.56~255.255.58.255",
		"23.135.167.228~255.0.0.0",
		"204.58.47.149~255.0.0.0",
		"113.33.181.46~255.255.255.0",
		"73.245.52.119~255.255.154.0",
		"23.214.47.71~255.0.0.0",
		"139.124.188.91~255.255.255.100",
		"142.94.192.197~255.0.0.0",
		"53.173.252.202~255.0.0.0",
		"127.201.56.50~255.255.111.255",
		"118.251.84.111~255.0.0.0",
		"130.27.73.170~255.0.0.0",
		"253.237.54.56~255.86.0.0",
		"64.189.222.111~255.255.255.139",
		"148.77.44.147~255.0.0.0",
		"59.213.5.253~255.255.0.0",
		"3.52.119.131~255.255.0.0",
		"213.208.164.145~255.255.0.0",
		"24.22.21.206~255.255.90.255",
		"89.43.34.31~255.0.0.0",
		"9.64.214.75~255.0.0.0",
		"71.183.242.53~255.255.0.0",
		"119.152.129.100~255.0.0.0",
		"38.187.119.201~255.0.0.0",
		"73.81.221.180~255.255.255.255",
		"73.198.13.199~255.0.15.0",
		"99.42.142.145~255.255.255.0",
		"196.121.115.160~255.0.0.0",
		"226.30.29.206~255.0.0.0",
		"244.248.31.171~255.255.255.255",
		"59.116.159.246~255.0.0.0",
		"121.124.37.157~255.0.0.226",
		"103.42.94.71~255.255.0.0",
		"125.88.217.249~255.255.74.255",
		"73.44.250.101~255.255.255.0",
	} {
		// line, _, _ := reader.ReadLine()
		// if len(line) == 0 {
		// 	break
		// }
		arr := strings.Split(line, "~")
		if !judgeIp(arr[1]) || !isMask(arr[1]) {
			f += 1
			A = A + "\n" + arr[0] + " " + arr[1]
			continue
		}
		if !judgeIp(arr[0]) {
			f += 1
			A = A + "\n" + arr[0] + " " + arr[1]
			continue
		}
		ip := arr[0]
		if isPrivate(ip) {
			g += 1
		}
		first, _ := strconv.Atoi(ip[:strings.Index(ip, ".")])
		if first > 0 && first < 127 {
			a += 1
		} else if first > 127 && first < 192 {
			b += 1
		} else if first > 191 && first < 224 {
			c += 1
		} else if first > 223 && first < 240 {
			d += 1
		} else if first > 239 && first < 256 {
			e += 1
		}
	}
	fmt.Printf("%d %d %d %d %d %d %d\n", a, b, c, d, e, f, g)
	fmt.Println(A)
	fmt.Println()
}
