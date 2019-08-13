package botupd

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"wkndbot/cnf"
	"wkndbot/service/botcmd"
	"wkndbot/service/security"
	"wkndbot/service/syscmd"
)

func Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil || !update.Message.IsCommand() || !security.IsAllowed(&update) {
		return
	}
	conf := cnf.Get()
	var msg tgbotapi.MessageConfig
	command := update.Message.Command()
	switch command {
	case "list":
		msg = botcmd.List(update)
	default:
		if len(conf.Commands[command].Command) == 0 {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
		} else {
			stdout := syscmd.Exec(conf.Commands[command])
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, string(stdout))
		}
	}
	log.Println(msg.Text)
	bot.Send(msg)
}
