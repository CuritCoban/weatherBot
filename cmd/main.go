package main

import (
	"fmt"
	"log"
	"strings"
	"weatherBot/config"
	"weatherBot/internal/handler"
	repo "weatherBot/internal/repository"
	weather "weatherBot/internal/usecases"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	defer func() {
		recover()
		log.Println("Fatal error")
	}()

	TOKEN, WEATHER_KEY, GORM_KEY := config.InitConfig()

	//Создание бота
	bot, err := botapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panicf("Init Bot error: %v", err)
	}
	log.Println("Bot init")

	//Создание канала обновлений
	updateConfig := botapi.NewUpdate(0)
	ch := bot.GetUpdatesChan(updateConfig)

	owApi := weather.ApiKey{Key: WEATHER_KEY}
	owApi.GetCoordinate("Moscow")

	tgbot := handler.TgBot{Bot: bot}
	db := repo.DataBase{Postgres: config.InitDB(GORM_KEY)}

	for update := range ch {
		if update.Message == nil {
			continue
		}
		//chatID := update.Message.Chat.ID
		msgText := update.Message.Text
		fmt.Printf("\nUser enter: %s", msgText)

		//Первый запуск
		if update.Message.Text == "/start" {
			tgbot.StartBot(update)
			continue
		}

		//Выбор города
		cityName := strings.Split(msgText, " ")
		if cityName[0] == "/selectCity" && len(cityName) > 1 {
			tgbot.SelectCity(update, db, cityName[1], owApi)
			continue
		}

		//Получение погоды в выбранном городе
		if msgText == "/weather" {
			tgbot.GetWeather(update, db, owApi)
			continue
		}

		//Если команда не найдена
		tgbot.EndBot(update)
	}
}

/*if coord.Lat == 0 && coord.Lon == 0 {
		bot.Send(botapi.NewMessage(chatID,
			fmt.Sprintf("Города '%s' не существует", msgText)))
		continue
	}
}
temp := owApi.GetTemperature(coord)
tempStr := fmt.Sprintf("%0.1f", temp)
bot.Send(botapi.NewMessage(chatID,
	fmt.Sprintf("Температура в %s: %s°C", msgText, tempStr)))*/
