package openWeatherMap

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"weatherBot/internal/models"
)

type ApiKey struct {
	Key string
}

func (k ApiKey) GetCoordinate(city string) models.Coordinate {
	//Получение координат города
	url := "http://api.openweathermap.org/geo/1.0/direct?q=" + city + "&limit=5&appid=" + k.Key
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		slog.Error("Coordinate", "Parse error", err)
		return models.Coordinate{}
	}

	//Распаковка json координат
	coordinate := models.CoordinateSlice{}
	err = json.NewDecoder(resp.Body).Decode(&coordinate)
	if err != nil {
		slog.Error("Coordinate", "Decode error", err)
		return models.Coordinate{}
	}
	if len(coordinate) == 0 {
		slog.Error("Coordinate is empty")
		return models.Coordinate{}
	}

	retCoordinate := coordinate[0]
	return retCoordinate
}

func (k ApiKey) GetTemperature(c models.Coordinate) float64 {
	w := models.Weather{}

	//Получение данных погоды по координатам
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric",
		c.Lat, c.Lon, k.Key)
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Temp", "Parse error", err)
		return -1000
	}

	//Распаковка json данных погоды
	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		slog.Error("Temp", "Decode error", err)
		return -1000
	}

	///fmt.Println("Weather in city: ", w.Main.Temp)
	return w.Main.Temp
}
