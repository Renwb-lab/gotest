package main

import (
	"fmt"
	"sync"
)

func process(f string) (string, int) {
	return f, len(f)
}
func main() {
	fileNames := []string{"1", "2", "3", "4", "1234"}
	ch := make(chan int, len(fileNames))

	var wg sync.WaitGroup
	for _, f := range fileNames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			_, l := process(f)
			ch <- l
		}(f)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var total int
	for s := range ch {
		total += s
	}
	fmt.Println(total)
}
