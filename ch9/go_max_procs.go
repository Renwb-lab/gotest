package main

import "fmt"

func main227() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}

}
