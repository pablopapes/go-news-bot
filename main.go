package main

func main() {

	infobae := &Infobae{url: "https://www.infobae.com"}
	lanacion := &LaNacion{url: "https://www.lanacion.com.ar"}
	clarin := &Clarin{url: "https://www.clarin.com"}

	feedList := []NewsFeed{infobae, lanacion, clarin}

	newsFeed := &initFeed{}

	telegramIntance := &TelegramBot{}
	bot := telegramIntance.CreateBotIntance()

	for _, feed := range feedList {
		newsFeed.setFeed(feed)
		newsFeed.ScraperNews(telegramIntance, bot)
	}
}
