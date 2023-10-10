package main

import (
	"fmt"
	"time"

	pool "gopkg.in/go-playground/pool.v3"
)

func SendMail(int int) pool.WorkFunc {
	fn := func(wu pool.WorkUnit) (interface{}, error) {
		// sleep 1s 模拟发邮件过程
		time.Sleep(time.Second * 1)
		// 模拟异常任务需要取消
		if int == 17 {
			wu.Cancel()
		}
		if wu.IsCancelled() {
			return false, nil
		}
		fmt.Println("send to", int)
		return true, nil
	}
	return fn
}

func main140() {
	// 初始化groutine数量为20的pool
	p := pool.NewLimited(20)
	defer p.Close()
	batch := p.Batch()
	// 设置一个批量任务的过期超时时间
	t := time.After(10 * time.Second)
	go func() {
		for i := 0; i < 100; i++ {
			batch.Queue(SendMail(i)) // 往批量任务中添加workFunc任务
			fmt.Println("add job", i)
		}
		// 通知批量任务不再接受新的workFunc, 如果添加完workFunc不执行改方法的话将导致取结果集时done channel一直阻塞
		batch.QueueComplete()
	}()
	// // 获取批量任务结果集, 因为 batch.Results 中要close results channel 所以不能将其放在LOOP中执行
	r := batch.Results()
LOOP:
	for {
		select {
		case <-t:
			// 超时通知
			fmt.Println("超时通知")
			break LOOP
		case email, ok := <-r:
			// 读取结果集
			if ok {
				if err := email.Error(); err != nil {
					fmt.Println("读取结果集错误，error info:", err.Error())
				}
				fmt.Println("结果集:", email.Value())
			} else {
				fmt.Println("finish")
				break LOOP
			}
		}
	}
}
