package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gocolly/colly"
)

type Infobae struct {
	url string
}

func (i *Infobae) CollectNews(telegramIntance *TelegramBot, telegramBot *tgbotapi.BotAPI) {
	collector := colly.NewCollector()
	database := &db{}
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
				article.url = i.url + e.ChildAttr(".headline-link", "href")
				article.feed = "🟡 Infobae"
				fmt.Println(article.title)
				fmt.Println(article.url)

				if !database.checkData(article) {
					telegramIntance.sendMessage(article, telegramBot)
					database.saveData(article)
				}
			}
		})
	})
	collector.Visit(i.url)
}
