package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// map to cache URLs
type urlCache struct {
	urlMap map[string]string
	mux    sync.Mutex
}

func (u *urlCache) Add(url string, body string) {
	u.mux.Lock()
	u.urlMap[url] = body
	u.mux.Unlock()
}

func (u *urlCache) Hit(url string) bool {
	u.mux.Lock()
	defer u.mux.Unlock()
	
	_, ok := u.urlMap[url]

	return ok
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *urlCache) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	if depth <= 0 {
		return
	}
	
	// No need to fetch URL if we already fetched it before
	if cache.Hit(url) {
		return	
	} else {
		cache.Add(url, body)
		body, urls, err := fetcher.Fetch(url)
		
		if err != nil {
			fmt.Println(err)
			return
		}
				
		fmt.Printf("found: %s %q\n", url, body)
		
		for _, u := range urls {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, cache)
		}
	}
	return
}

// Wait group to handle first call to Crawl
var wg sync.WaitGroup

func main() {
	// Create URL cache
	u := urlCache{urlMap: make(map[string]string)}
	cache := &u
	
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, cache)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
