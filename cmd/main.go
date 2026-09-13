package main

import (
	"fmt"
	"log"
	"weatherBot/internal/config"
	"weatherBot/internal/handler"
	weather "weatherBot/internal/usecases"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	defer func() {
		recover()
		log.Println("Fatal error")
	}()

	TOKEN, WEATHER_KEY := config.InitConfig()

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
	tgbot := handler.TgBot{Bot: bot}

	for update := range ch {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		msgText := update.Message.Text
		fmt.Printf("\nUser enter: %s", msgText)

		//Обработка первого запуска; Отправка
		tgbot.StartBot(update)

		coord := owApi.GetCoordinate(msgText)
		//Обработка неверного ввода; Отправка
		if coord.Lat == 0 && coord.Lon == 0 {
			bot.Send(botapi.NewMessage(chatID,
				fmt.Sprintf("Города '%s' не существует", msgText)))
			continue
		}

		//Получение температуры и округление до .0; Отправка
		temp := owApi.GetTemperature(coord)
		tempStr := fmt.Sprintf("%0.1f", temp)
		bot.Send(botapi.NewMessage(chatID,
			fmt.Sprintf("Температура в %s: %s°C", msgText, tempStr)))
	}
}
