package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Article struct {
	title string
	url   string
}

type NewsFeed interface {
	CollectNews(telegramIntance *TelegramBot, telegramBot *tgbotapi.BotAPI)
}

type initFeed struct {
	newsFeed NewsFeed
}

func (i *initFeed) setFeed(newsFeed NewsFeed) {
	i.newsFeed = newsFeed
}

func (i *initFeed) ScraperNews(telegramIntance *TelegramBot, telegramBot *tgbotapi.BotAPI) {
	i.newsFeed.CollectNews(telegramIntance, telegramBot)
}
