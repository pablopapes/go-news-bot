package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Infobae struct {
	url string
}

func (i *Infobae) CollectNews() {
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
	collector.OnHTML(".three-elements-chain", func(e *colly.HTMLElement) {
		e.ForEach(".story-card-info", func(_ int, e *colly.HTMLElement) {
			if e.ChildText(".story-card-hl") != "" {
				article := Article{}
				article.title = e.ChildText(".story-card-hl")
				article.url = e.ChildAttr(".headline-link", "href")
				fmt.Println(article.title)
				fmt.Println(i.url + article.url)

				/*
					msg := tgbotapi.NewMessageToChannel("@news_argy", article.Title+"\n"+"👉 <a href='"+urlInfobae+article.Url+"'>Link a infobae</a>")
					msg.ParseMode = "HTML"
					msg.DisableWebPagePreview = true
					bot.Send(msg)
				*/

			}
		})
	})
	collector.Visit(i.url)
}
