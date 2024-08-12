package main

type Article struct {
	title string
	url   string
}

type initFeed struct {
	newsFeed NewsFeed
}

func (i *initFeed) setFeed(newsFeed NewsFeed) {
	i.newsFeed = newsFeed
}

func (i *initFeed) ScraperNews() {
	i.newsFeed.CollectNews()
}

type NewsFeed interface {
	CollectNews()
}
