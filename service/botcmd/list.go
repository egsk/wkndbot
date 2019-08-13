package botcmd

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"wkndbot/cnf"
)

func List(update tgbotapi.Update) tgbotapi.MessageConfig {
	var commands []string
	for command := range cnf.Get().Commands {
		commands = append(commands, command)
	}
	var btns []tgbotapi.KeyboardButton
	for i := range commands {
		btns = append(btns, tgbotapi.NewKeyboardButton("/" + commands[i]))
	}
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			btns...
		),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список доступных команд")
	msg.ReplyMarkup = keyboard

	return msg
}
