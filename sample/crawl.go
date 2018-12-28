package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)

	// 从命令行参数开始
	go func() {
		worklist <- os.Args[1:]
	}()

	// 并发爬取web
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
