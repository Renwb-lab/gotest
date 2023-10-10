package main

import (
	"fmt"
	"time"
)

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{
		f: f,
	}
}

func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	EntryChannel chan *Task
	WorkNum      int
	JobsChannel  chan *Task
}

func NewPool(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		WorkNum:      cap,
		JobsChannel:  make(chan *Task),
	}
}

func (p *Pool) Worker(workID int) {
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("Worker ID", workID, " finished.")
		time.Sleep(time.Second)
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.WorkNum; i += 1 {
		fmt.Println("开启固定数量的Worker:", i)
		go p.Worker(i)
	}

	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}

	close(p.EntryChannel)
	close(p.JobsChannel)
}

func main138() {
	t := NewTask(func() error {
		fmt.Println("创建一个Task:", time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})

	p := NewPool(5)

	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	p.Run()
}
