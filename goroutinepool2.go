package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task2 struct {
	f func() error
}

func NewTask2(f func() error) *Task2 {
	return &Task2{
		f: f,
	}
}

func (t *Task2) Execute() {
	t.f()
}

type Pool2 struct {
	Channel chan *Task2
	Cap     int
	Wg      *sync.WaitGroup
}

func NewPool2(cap int) *Pool2 {
	return &Pool2{
		Channel: make(chan *Task2),
		Cap:     cap,
		Wg:      &sync.WaitGroup{},
	}
}

func (p *Pool2) Worker(ctx context.Context, workID int) {
	for {
		select {
		case <-ctx.Done():
			p.Wg.Done()
			fmt.Println("Worker ID", workID, " exit.")
			return
		case task := <-p.Channel:
			task.Execute()
			fmt.Println("Worker ID", workID, " finished.")
			time.Sleep(time.Second)
		}
	}
}

func (p *Pool2) Run(ctx context.Context) {
	for i := 0; i < p.Cap; i += 1 {
		fmt.Println("开启固定数量的Worker:", i)
		p.Wg.Add(1)
		go p.Worker(ctx, i)
	}
	p.Wg.Wait()
	fmt.Println("任务结束")
}

func main139() {
	t := NewTask2(func() error {
		fmt.Println("创建一个Task:", time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})

	cxt, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	p := NewPool2(5)

	go func() {
		for {
			select {
			case <-cxt.Done():
				return
			case p.Channel <- t:
			}
		}
	}()

	p.Run(cxt)
}
