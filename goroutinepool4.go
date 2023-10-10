package main

import (
	"fmt"
	"sync"
	"time"

	ants "github.com/panjf2000/ants"
)

// 任务
func sendMail(i int, wg *sync.WaitGroup) func() {
	var cnt int
	return func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("send mail to ", i, "cnt: ", cnt)
			cnt++
			if cnt > 5 {
				fmt.Println("退出协程ID:", i)
				break
			}
		}
		wg.Done()
	}
}

func main150() {
	wg := sync.WaitGroup{}

	//申请一个协程池对象
	pool, _ := ants.NewPool(2)

	//关闭协程池
	defer pool.Release()

	// 向pool提交任务
	for i := 1; i <= 5; i++ {
		pool.Submit(sendMail(i, &wg))
		wg.Add(1)
	}
	wg.Wait()
}
