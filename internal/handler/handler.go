package handler

import (
	"log"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TgBot struct {
	Bot *botapi.BotAPI
}
type DataBase struct {
	Postgres *gorm.DB
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

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=3429 dbname=postgres port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("DataBase init")
	return db
}
