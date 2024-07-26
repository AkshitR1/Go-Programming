package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link found:", link)
		// Visit the link found
		e.Request.Visit(link)
	})

	// Start scraping on a website
	c.Visit("http://example.com")
}
