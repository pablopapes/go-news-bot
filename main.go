package main

import (
	"fmt"
	"os"

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

func main() {

	/*
		bot, err := tgbotapi.NewBotAPI(goDotEnvVariable("API_KEY"))
		if err != nil {
			panic(err)

		bot.Debug = false
		}
	*/

	infobae := &Infobae{url: "https://www.infobae.com"}
	lanacion := &LaNacion{url: "https://www.lanacion.com.ar"}
	clarin := &Clarin{url: "https://www.clarin.com"}

	feedList := []NewsFeed{infobae, lanacion, clarin}

	newsFeed := &initFeed{}

	for _, feed := range feedList {
		newsFeed.setFeed(feed)
		newsFeed.ScraperNews()
	}
}
