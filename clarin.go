package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Clarin struct {
	url string
}

func (l *Clarin) CollectNews() {

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
	collector.OnHTML(".Container2Pisos", func(e *colly.HTMLElement) {
		e.ForEach("article", func(_ int, e *colly.HTMLElement) {
			if e.ChildText(".title") != "" {
				article := Article{}
				article.title = e.ChildText(".title")
				article.url = e.ChildAttr("a", "href")
				fmt.Println(article.title)
				fmt.Println(article.url)
				/*
					msg := tgbotapi.NewMessageToChannel("@news_argy", article.title+"\n"+"👉 <a href='"+l.url+article.url+"'>Link a La Nación</a>")
					msg.ParseMode = "HTML"
					msg.DisableWebPagePreview = true
					bot.Send(msg)
				*/
			}
		})
	})
	collector.Visit(l.url)
}
