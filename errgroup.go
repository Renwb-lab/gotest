package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main071034() {
	var g errgroup.Group
	var urls = []string{
		"http://blog.csdn.net/",
		"http://www.baidu.com/",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err.Error())
	}
}
