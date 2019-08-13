package security

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"wkndbot/cnf"
)

func IsAllowed(update *tgbotapi.Update) bool {
	conf := cnf.Get()
	currentUser := update.Message.From.UserName
	isAllowed := false
	for _, user := range conf.Users.Allowed {
		if currentUser == user {
			isAllowed = true
		}
	}

	return isAllowed
}
