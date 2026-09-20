package main

import (
	"log"
	"log/slog"
	"weatherBot/config"
	"weatherBot/internal/handler"
	repo "weatherBot/internal/repository"
	weather "weatherBot/internal/usecases"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	defer func() {
		recover()
		slog.Error("Fatal error")
	}()
	//Загрузка переменных окружения
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
	tgbot := handler.TgBot{Bot: bot}
	db := repo.DataBase{Postgres: config.InitDB(GORM_KEY)}

	//Переменная ожидания
	wait := make(map[int64]bool)

	//Создание кнопок
	button := botapi.NewReplyKeyboard(
		botapi.NewKeyboardButtonRow(
			botapi.NewKeyboardButton("Город"),
			botapi.NewKeyboardButton("Погода"),
		),
	)

	//Цикл обработки сообщений
	for update := range ch {
		chatID := update.Message.Chat.ID
		msgText := update.Message.Text
		if update.Message == nil {
			continue
		}

		switch {
		//Первый запуск /start
		case msgText == "/start":
			msg := botapi.NewMessage(chatID,
				"Команда 'Город' чтобы выбрать город где нужно узнать погоду.\n"+
					"Команда 'Погода' чтобы показать погоду в выбранном городе")
			bot.Send(msg)

		//Выбор города
		case msgText == "Город" || wait[chatID] == true:
			if wait[chatID] == true {
				tgbot.SelectCity(update, db, msgText, owApi)
				wait[chatID] = false
				break
			}
			wait[chatID] = true
			bot.Send(botapi.NewMessage(chatID, "Введи город: "))

		//Отправка погоды в выбранном городе
		case msgText == "Погода":
			tgbot.GetWeather(update, db, owApi)

		//Появление кнопок
		default:
			msg := botapi.NewMessage(chatID, "Выбери действие")
			msg.ReplyMarkup = button
			bot.Send(msg)
		}
	}
}
