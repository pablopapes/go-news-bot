package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

type Article struct {
	Title string
	Url   string
}

func main() {
	args := os.Args
	urlInfobae := args[1]
	urlLaNacion := args[2]
	collector := colly.NewCollector()

	bot, err := tgbotapi.NewBotAPI(goDotEnvVariable("API_KEY"))
	if err != nil {
		panic(err)
	}

	bot.Debug = false

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
			if e.ChildText(".story-card-hl") != "" {
				article := Article{}
				article.Title = e.ChildText(".story-card-hl")
				article.Url = e.ChildAttr(".headline-link", "href")
				fmt.Println(article.Title)
				fmt.Println(urlInfobae + article.Url)

				msg := tgbotapi.NewMessageToChannel("@news_argy", article.Title+"\n"+"👉 <a href='"+urlInfobae+article.Url+"'>Link a infobae</a>")
				msg.ParseMode = "HTML"
				msg.DisableWebPagePreview = true
				bot.Send(msg)

			}
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

			msg := tgbotapi.NewMessageToChannel("@news_argy", article.Title+"\n"+"👉 <a href='"+urlInfobae+article.Url+"'>Link a La Nación</a>")
			msg.ParseMode = "HTML"
			msg.DisableWebPagePreview = true
			bot.Send(msg)
		})
	})

	collector.Visit(urlInfobae)
	collector.Visit(urlLaNacion)

}
