package bot

import (
	"fmt"
	"log"
	"terceirapontebot/crawler"

	"gopkg.in/telegram-bot-api.v4"
)

func sendMessage(c chan int64, chatID int64, path string, bot *tgbotapi.BotAPI) {
	for {
		bytes := <-c
		msg := tgbotapi.NewPhotoUpload(chatID, path)
		log.Printf("%v bytes saved", bytes)

		_, err := bot.Send(msg)
		if err != nil {
			panic(err)
		}
	}
}

func parseUpdate(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	pageURL := "http://www.rodosol.com.br/"
	chatID := update.Message.Chat.ID

	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.Command(), update.Message.Text)
	if update.Message.Command() != "now" {
		return
	}

	bot.Send(tgbotapi.NewMessage(chatID, "Getting images..."))

	links := crawler.GetImgLinks(pageURL)
	for i := 0; i < len(links); i++ {
		c := make(chan int64)

		path := fmt.Sprintf("img_%v.jpg", i)

		go crawler.SaveImage(c, links[i].Attr[0].Val, path)
		go sendMessage(c, chatID, path, bot)
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
