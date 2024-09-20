package main

import (
	"os"
)

// // forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// // pre和post都是可选的。
// // 遍历孩子结点之前，pre被调用
// // 遍历孩子结点之后，post被调用
// func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
// 	if pre != nil {
// 		pre(n)
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		forEachNode(c, pre, post)
// 	}
// 	if post != nil {
// 		post(n)
// 	}
// }

// // Extract makes an HTTP GET request to the specified URL, parses
// // the response as HTML, and returns the links in the HTML document.
// func Extract(url string) ([]string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		resp.Body.Close()
// 		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
// 	}
// 	doc, err := html.Parse(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
// 	}
// 	var links []string
// 	visitNode := func(n *html.Node) {
// 		if n.Type == html.ElementNode && n.Data == "a" {
// 			for _, a := range n.Attr {
// 				if a.Key != "href" {
// 					continue
// 				}
// 				link, err := resp.Request.URL.Parse(a.Val)
// 				if err != nil {
// 					continue // ignore bad URLs
// 				}
// 				links = append(links, link.String())
// 			}
// 		}
// 	}
// 	forEachNode(doc, visitNode, nil)
// 	return links, nil
// }

// func crawl(url string) []string {
// 	fmt.Println(url)
// 	list, err := Extract(url)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	return list
// }

func main234() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
