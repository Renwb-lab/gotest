package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main071014() {
	rand.Seed(time.Now().UnixNano())
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	var wg sync.WaitGroup
	// It must be a buffered channel.
	toStop := make(chan string, 1)

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		fmt.Println(stoppedBy)
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i += 1 {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select {
					case toStop <- "sender#" + id: //退出
					default:
					}
					return
				}
				select {
				case <-stopCh: // 检查是否终止
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i += 1 {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			for {
				select {
				case <-stopCh: // 检查是否终止
					return
				case value := <-dataCh:
					// fmt.Println(value)
					if value == Max-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
				}
			}
		}(strconv.Itoa(i))
	}
	wg.Wait()
}
