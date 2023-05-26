package main

import (
	"fmt"
)

// 向一个未初始化的ch进行对数据的话，返回结果为默认值。
func main96() {
	var ch chan int
	c, ok := <-ch
	fmt.Println(ok)
	fmt.Println(c)
	fmt.Println("end")
}
