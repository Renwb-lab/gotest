package main

import (
	"os"
	"runtime/trace"
)

func main323() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	ch := make(chan string)
	go func() {
		ch <- "SEND"
	}()
	<-ch
}
