package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"wkndbot/cnf"
	"wkndbot/service/botupd"
)

func main() {
	conf := cnf.Get()
	bot, err := tgbotapi.NewBotAPI(conf.Credentials.Token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		botupd.Handle(bot, update)
	}
}
