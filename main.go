package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Article struct {
	Title string
	Url   string
}

func main() {
	args := os.Args
	urlInfobae := args[1]
	urlLaNacion := args[2]
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})

	// infobae
	collector.OnHTML(".three-elements-chain", func(e *colly.HTMLElement) {
		e.ForEach(".story-card-info", func(_ int, e *colly.HTMLElement) {
			article := Article{}
			article.Title = e.ChildText(".story-card-hl")
			article.Url = e.ChildAttr("a", "href")
			fmt.Println(article.Title)
			fmt.Println(urlInfobae + article.Url)
		})
	})

	// la nacion
	collector.OnHTML(".ln-opening-container", func(e *colly.HTMLElement) {
		e.ForEach(".ln-card", func(_ int, e *colly.HTMLElement) {
			article := Article{}
			article.Title = e.ChildText(".title")
			article.Url = e.ChildAttr("section", "href")
			fmt.Println(article.Title)
			fmt.Println(urlLaNacion + article.Url)
		})
	})

	collector.Visit(urlInfobae)
	collector.Visit(urlLaNacion)
}
