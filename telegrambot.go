package main

import (
	"fmt"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type TelegramBot struct {
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func (t *TelegramBot) CreateBotIntance() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(goDotEnvVariable("API_KEY"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	return bot
}

func (t *TelegramBot) sendMessage(article Article, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessageToChannel("@news_argy", article.feed+"\n\n"+article.title+"\n\n"+"<a href='"+article.url+"'>Link</a>")
	msg.ParseMode = "HTML"
	msg.DisableWebPagePreview = true
	bot.Send(msg)
	fmt.Println("Waiting 10 seconds to start the next message")
	time.Sleep(10 * time.Second)
}
