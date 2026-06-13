package main

import (
	"fmt"
	"sync"
)

func WebScrapper(sites string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fetching", sites)
	fmt.Println("Done fetching", sites)
}

func main() {
	var wg sync.WaitGroup

	Sites := []string{
		"google.com",
		"github.com",
		"youtube.com",
		"golang.org",
	}
	for i := range Sites {
		wg.Add(1)
		go WebScrapper(Sites[i], &wg)

	}
	wg.Wait()
	fmt.Println("All websites fetched")

}
