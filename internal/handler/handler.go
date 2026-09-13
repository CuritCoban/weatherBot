package handler

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	Bot *botapi.BotAPI
}

func (b *TgBot) StartBot(update botapi.Update) {
	if update.Message.Text == "/start" {
		b.Bot.Send(botapi.NewMessage(
			update.Message.Chat.ID, "Hi"))
	}
}

func (b *TgBot) SelectCity(update botapi.Update) {
	if update.Message.Text == "/selectCity" {

	}
}
