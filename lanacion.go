package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gocolly/colly"
)

type LaNacion struct {
	url string
}

func (l *LaNacion) CollectNews(telegramIntance *TelegramBot, telegramBot *tgbotapi.BotAPI) {

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
	collector.OnHTML(".ln-opening-container", func(e *colly.HTMLElement) {
		e.ForEach(".ln-card", func(_ int, e *colly.HTMLElement) {
			article := Article{}
			article.title = e.ChildText(".title")
			article.url = l.url + e.ChildAttr("section", "href")
			article.feed = "🔵 La Nación"
			fmt.Println(article.title)
			fmt.Println(article.url)
			if !database.checkData(article) {
				telegramIntance.sendMessage(article, telegramBot)
				database.saveData(article)
			}
		})
	})
	collector.Visit(l.url)
}
