package main

import (
	"fmt"
	"log"
	"terceirapontebot/bot"

	"github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
)

func getToken() string {
	viper.SetConfigName("secret")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	return viper.Get("BOT_TOKEN").(string)
}

func main() {
	botAPI, err := tgbotapi.NewBotAPI(getToken())
	if err != nil {
		log.Panic("Can't authenticate: ", err)
	}

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	bot.Start(botAPI)
}
