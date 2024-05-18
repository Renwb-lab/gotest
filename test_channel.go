package main

import "fmt"

func main() {
	{
		var c chan int
		c <- 1
	}
	{
		var c chan int
		a := <-c
		fmt.Print(a)
	}

	{
		c := make(chan int)
		close(c)
		a := <-c
		fmt.Print(a)
	}
	{
		c := make(chan int)
		close(c)
		c <- 1
	}

	{
		c := make(chan int)
		close(c)
		close(c)
	}

}
