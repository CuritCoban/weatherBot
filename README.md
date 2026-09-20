# weatherBot
## Описание 
### Суть работы
1. Бот парсит координаты указанного города 
2. Добавляеет их в БД 
3. Достает их из БД и парсит температуру с [OpenWeatherMap](https://openweathermap.org/api)
4. Отправляет результат пользователю
#### Docker
Docker не работает. Что то запускается что то нет
## Стек
- Go
- PostgreSQL
- GORM
- Telegram Bot Api
- OpenWeatherMap Api
- Git
- Goose
- Docker 50/50
## Установка
1. git clone https://github.com/CuritCoban/weatherBot.git
2. cd weatherBot
3. go mod download
4. В weatherBot/cmd создать `.env` где будет:
```
TOKEN="токен бота"
WEATHER_KEY="ключ OpenWeatherMap"
POSTGRES_PASSWORD="пароль"
POSTGRES_DB="weatherbot"
GORM_KEY="host=localhost user=postgres password=пароль dbname=weatherbot port=9920 sslmode=disable TimeZone=Asia/Shanghai"

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres:пароль@localhost:9920/weatherbot
GOOSE_MIGRATION_DIR= целиком путь к /migrations
```
5. goose up
6. Готово