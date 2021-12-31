package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mmcdole/gofeed/rss"
)

func main() {
	file, err := os.Open("hackernews.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Feed Parser
	fp := rss.Parser{}
	feed, _ := fp.Parse(io.Reader(file))
	if err != nil {
		panic(err)
	}
	// feed.Items has a k:v result
	for _, v := range feed.Items {
		fmt.Printf(`Title: %sThread Link: %sContent: %s`, v.Title, v.Link, v.Comments)
	}
}
