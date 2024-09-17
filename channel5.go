package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const N = 5

var APIList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

func main217() {
	type JobInfo struct {
		name string
		ret  error
	}
	// 1. 创建一个channel，用于通知可以执行
	nextChan := make(chan string, 5)
	// 2. 创建一个channel，用于接受返回
	retChan := make(chan JobInfo, len(APIList))
	startChan := make(chan struct{})
	ctx, cancal := context.WithCancel(context.Background())
	defer cancal()
	var sg sync.WaitGroup
	// 4. 启动投放worker
	sg.Add(1)
	go func(ctx context.Context) {
		defer sg.Done()
		<-startChan
		for _, s := range APIList {
			select {
			case nextChan <- s:
				// fmt.Printf("Push %s\n", s)
			case <-ctx.Done():
				// fmt.Printf("Push error\n")
				return
			}
		}
		// fmt.Println("Push end 2")
	}(ctx)
	// 3. 启动worker
	for i := 0; i < N; i += 1 {
		sg.Add(1)
		go func(cxt context.Context, name string) {
			defer sg.Done()
			for {
				// fmt.Println(name, cap(nextChan))
				select {
				case s := <-nextChan:
					// fmt.Printf("Work %s\n", s)
					ret := callAPI(s)
					retChan <- JobInfo{name: s, ret: ret}
				case <-ctx.Done():
					// fmt.Printf("Work Done 1\n")
					return
				}
			}
		}(ctx, strconv.Itoa(i))
	}
	// 5. 主线程接受结果
	idx := 0
	startChan <- struct{}{}
	for r := range retChan {
		fmt.Printf("Receive %#v\n", r)
		if r.ret != nil || idx == len(APIList)-1 {
			if r.ret != nil {
				fmt.Println(r.name, r.ret)
			}
			cancal()
			break
		}
		idx += 1
	}

	sg.Wait()
}

func callAPI(a string) error {
	time.Sleep(time.Second)
	rand.Seed(time.Now().UnixNano())
	n := rand.Int31n(10)
	if n < 2 {
		return fmt.Errorf("%s error, rate: %d", a, n)
	}
	return nil
}
