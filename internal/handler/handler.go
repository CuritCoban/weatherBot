package handler

import (
	"fmt"
	"log/slog"
	"weatherBot/internal/models"
	repo "weatherBot/internal/repository"
	weather "weatherBot/internal/usecases"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	Bot *botapi.BotAPI
}

// Выбор города
func (b *TgBot) SelectCity(update botapi.Update, db repo.DataBase, cityName string, owApi weather.ApiKey) {
	//Обработки ошибки ввода
	coord := owApi.GetCoordinate(cityName)
	if coord.Lat == 0 && coord.Lon == 0 {
		slog.Warn("Incorrect input", "cityName", cityName)

		b.Bot.Send(botapi.NewMessage(update.Message.Chat.ID,
			fmt.Sprintf("Города '%s' не существует", cityName)))
		return
	}

	//Заполнение структуры нужными данными
	city := models.WeatherDB{
		ChatID: update.Message.Chat.ID,
		City:   cityName,
		Temp:   0.00,
		Lon:    coord.Lon,
		Lat:    coord.Lat,
	}

	//Поиск в таблице выбранного города
	result := db.Postgres.Table("weather").Where("chat_id = ?",
		update.Message.Chat.ID).Save(city)

	//Если город не найден
	if result.Error != nil {
		slog.Error("Table", "Not found error", result.Error)

		//Создается новая таблица
		result = db.Postgres.Table("weather").Create(city)

		//Если создание не успешно выход из функции
		if result.Error != nil {
			slog.Error("Table", "Create error", result.Error)

			b.Bot.Send(botapi.NewMessage(
				update.Message.Chat.ID, "Город не выбран: "+cityName))
			return
		}
		slog.Info("Table create success", "ChatID", city.ChatID, "City", city.City)
	}

	b.Bot.Send(botapi.NewMessage(
		update.Message.Chat.ID, "Выбран город: "+cityName))
	slog.Info("Table found success", "ChatID", city.ChatID, "City", city.City)
}

// Отправка погоды
func (b *TgBot) GetWeather(update botapi.Update, db repo.DataBase, owApi weather.ApiKey) {
	weather := models.WeatherDB{}
	result := db.Postgres.Table("weather").Where("chat_id = ?",
		update.Message.Chat.ID).Find(&weather)

	if result.Error != nil || result.RowsAffected == 0 {
		slog.Error("Find data in DB", "Error", result.Error, "RowsAffected", result.RowsAffected)
		return
	}

	switch weather.Temp {

	case 0.00:
		weather.Temp = owApi.GetTemperature(models.Coordinate{Lon: weather.Lon, Lat: weather.Lat})
		if weather.Temp == -1000 {
			slog.Error("GetTemperature error")
			b.Bot.Send(botapi.NewMessage(update.Message.Chat.ID, "Server error"))
			return
		}
		result = db.Postgres.Table("weather").Where("chat_id = ?",
			update.Message.Chat.ID).Update("temp", weather.Temp)
		if result.Error != nil {
			slog.Error("Table", "Update error", result.Error)
			return
		}

		slog.Info("Send", "Temp", weather.Temp, "City", weather.City)
		b.Bot.Send(botapi.NewMessage(
			update.Message.Chat.ID, fmt.Sprintf("Температура в городе %s: %.0f°C",
				weather.City, weather.Temp)))
		return

	default:
		slog.Info("Send", "Temp", weather.Temp, "City", weather.City)
		b.Bot.Send(botapi.NewMessage(
			update.Message.Chat.ID, fmt.Sprintf("Температура в городе %s: %.0f°C",
				weather.City, weather.Temp)))
		return
	}
}
