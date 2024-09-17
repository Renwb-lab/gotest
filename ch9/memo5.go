package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
			// fmt.Println("not hit")
		} else {
			// fmt.Println("hit")
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string {
	return []string{
		"https://www.qq.com",
		"https://godoc.org",
		"https://www.baidu.com",
		"http://gopl.io",
		"https://www.qq.com",
		"https://godoc.org",
		"https://www.baidu.com",
		"http://gopl.io",
	}
}
func main226() {
	m := New(httpGetBody)
	var n sync.WaitGroup
	for i := 0; i < 5; i += 1 {
		for _, url := range incomingURLs() {
			n.Add(1)
			go func(url string) {
				defer n.Done()

				start := time.Now()
				value, err := m.Get(url)
				if err != nil {
					log.Print(err)
				}
				fmt.Printf("%s, %s, %d bytes\n",
					url, time.Since(start), len(value.([]byte)))
			}(url)
		}
		n.Wait()
		fmt.Println("==============================")
	}

}
