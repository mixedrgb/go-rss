package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/mmcdole/gofeed/rss"
)

func main() {
	url := "https://www.digitalocean.com/blog/rss/"
	// open a test file
	//file, err := os.Open("rss-test-hackernews.txt")

	// use the http.Get() method, returns response and error
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//defer file.Close()
	defer resp.Body.Close()

	// Feed Parser
	fp := rss.Parser{}

	feed, err := fp.Parse(io.Reader(resp.Body))
	//feed, _ := fp.Parse(io.Reader(file))

	if err != nil {
		log.Fatalf("Error with: %v\n", err)
	}
	c := color.New(color.FgCyan).Add(color.BgHiMagenta)
	// feed.Items has a k:v result
	for _, v := range feed.Items {
		fmt.Printf("\nDescription:")
		c.Print(v.Description)
	}
}
