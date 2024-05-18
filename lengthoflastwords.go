package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(s string) int {
	b := []byte(s)
	i := len(b) - 1
	for i >= 0 && b[i] == ' ' {
		i -= 1
	}
	r := 0
	for i >= 0 && b[i] != ' ' {
		r = r + 1
		i -= 1
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		b, _, _ := reader.ReadLine()
		if len(b) == 0 {
			return
		}
		fmt.Println(string(b))
		r := f(string(b))
		fmt.Printf("%d\n", r)
	}
}
