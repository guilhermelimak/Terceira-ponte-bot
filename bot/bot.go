package bot

import (
	"fmt"
	"log"
	"terceirapontebot/crawler"

	"gopkg.in/telegram-bot-api.v4"
)

func parseUpdate(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	pageURL := "http://www.rodosol.com.br/blog/categoria/terceira-ponte"

	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.Command(), update.Message.Text)
	if update.Message.Command() != "now" {
		return
	}

	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Getting images..."))

	links := crawler.GetImgLinks(pageURL)
	for i := 0; i < len(links); i++ {
		path := fmt.Sprintf("img_%v.jpg", i)

		crawler.SaveImage(links[i].Attr[0].Val, path)
		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, path)

		_, err := bot.Send(msg)
		if err != nil {
			panic(err)
		}
	}
}

// Start : initialize bot connection
func Start(bot *tgbotapi.BotAPI) {
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
