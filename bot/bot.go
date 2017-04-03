package bot

import (
	"log"
	"os"
	"terceirapontebot/crawler"

	"gopkg.in/telegram-bot-api.v4"
)

func parseUpdate(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.Command(), update.Message.Text)
	if update.Message.Command() != "now" {
		return
	}

	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Getting images"))

	links := crawler.GetPage("http://www.rodosol.com.br/blog/categoria/terceira-ponte")
	for i := 0; i < len(links); i++ {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, links[i].Attr[0].Val))
	}
}

// Start : initialize bot connection
func Start() {
	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		parseUpdate(&update, bot)
	}
}
