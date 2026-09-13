package handler

import (
	"fmt"
	"log"
	"weatherBot/internal/models"
	repo "weatherBot/internal/repository"
	weather "weatherBot/internal/usecases"

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

func (b *TgBot) SelectCity(update botapi.Update, db repo.DataBase, cityName string, owApi weather.ApiKey) {
	coord := owApi.GetCoordinate(cityName)
	if coord.Lat == 0 && coord.Lon == 0 {
		b.Bot.Send(botapi.NewMessage(update.Message.Chat.ID,
			fmt.Sprintf("Города '%s' не существует", cityName)))
		return
	}

	city := models.WeatherDB{
		ChatID:    update.Message.Chat.ID,
		City:      cityName,
		Temp:      0.00,
		Lon:       coord.Lon,
		Lat:       coord.Lat,
		CreatedAt: update.Message.Time(),
	}

	result := db.Postgres.Table("weather").Where("chat_id = ?, city = ?",
		update.Message.Chat.ID, cityName).First(&city)
	if result.Error != nil {
		log.Println("Error creating table: ", result.Error)

		result = db.Postgres.Table("weather").Create(&city)
		if result.Error != nil {
			b.Bot.Send(botapi.NewMessage(
				update.Message.Chat.ID, "Город не выбран: "+cityName))
			log.Println("Error creating table: ", result.Error)
			return
		}
	}

	b.Bot.Send(botapi.NewMessage(
		update.Message.Chat.ID, "Выбран город: "+cityName))
	log.Println("City created success: ", city)
}

func (b *TgBot) GetWeather(update botapi.Update, db repo.DataBase, owApi weather.ApiKey) {
	weather := models.WeatherDB{}
	result := db.Postgres.Table("weather").Where("chat_id = ?",
		update.Message.Chat.ID).First(&weather)
	if result.Error != nil {
		log.Println("Error writing to 'weather': ", result.Error)
		return
	}

	switch weather.Temp {

	case 0.00:
		weather.Temp = owApi.GetTemperature(
			models.Coordinate{Lon: weather.Lon, Lat: weather.Lat})

		result = db.Postgres.Table("weather").Where("chat_id = ?",
			update.Message.Chat.ID).Update("temp", weather.Temp)
		if result.Error != nil {
			log.Println("Error update database: ", result.Error)
			return
		}
		b.Bot.Send(botapi.NewMessage(
			update.Message.Chat.ID, fmt.Sprintf("Температура в городе %s: %.0f°C",
				weather.City, weather.Temp)))
		return

	default:
		b.Bot.Send(botapi.NewMessage(
			update.Message.Chat.ID, fmt.Sprintf("Температура в городе %s: %.0f°C",
				weather.City, weather.Temp)))
		return
	}
}
