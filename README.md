# go-telegram-news-bot
Update Telegram channel with the latest news from your favorite news site.

## Usage
1. [Create a telegram bot](https://core.telegram.org/bots#how-do-i-create-a-bot)
2. Create a telegram channel and add the bot as an admin.
2. Add a new A record in your Cloudflare account with the name you want to update ( Proxy status: Proxied)
3. Clone this repository
4. edit your .env file with your telegram API KEY and the Telegram channel ID.
EXAMPLE:
```
API_KEY=YOUR_API_KEY
CHANNEL_ID=@YOUR_CHANNEL_ID
```

5. Compile the code
```
go build
```
6. Run the code
```
./newsbot
```
7. You can use a cron job to run the code every X minutes/hours/days
```
*/5 * * * * /path/to/newsbot
```
## Create your own news feed
1. Create the `your_news.go` file and add your own news scraper logic. ( see `infobae.go` for an example)
2. Add your news scraper to the `main.go` file