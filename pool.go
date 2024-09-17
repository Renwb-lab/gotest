package main

import (
	"fmt"
	"sync"
	"time"
)

type MyObject struct {
	Name string
	// 其他字段...
}

func main291() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new object")
			return &MyObject{Name: "Default"}
		},
	}

	var wg sync.WaitGroup
	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			obj := pool.Get().(*MyObject)
			fmt.Printf("Goroutine %d got object with name: %s\n", id, obj.Name)

			obj.Name = fmt.Sprintf("Object_%d", id)
			pool.Put(obj)
		}(i)
	}

	wg.Wait()
}
