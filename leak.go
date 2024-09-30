package main

func leak() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
}
