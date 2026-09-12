package openWeatherMap

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"weatherBot/internal/models"
)

type ApiKey struct {
	Key string
}

func (k ApiKey) GetCoordinate(city string) models.Coordinate {
	url := "http://api.openweathermap.org/geo/1.0/direct?q=" + city + "&limit=5&appid=" + k.Key
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Parse coordinate error: %v", err)
		return models.Coordinate{}
	}
	coordinate := models.CoordinateSlice{}
	err = json.NewDecoder(resp.Body).Decode(&coordinate)
	if err != nil {
		log.Printf("Decode coordinate error: %v", err)
	}
	if len(coordinate) == 0 {
		log.Println("coordinate is empty")
		return models.Coordinate{}
	}

	retCoordinate := coordinate[0]
	return retCoordinate
}

func (k ApiKey) GetTemperature(c models.Coordinate) float64 {
	w := models.Weather{}
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric",
		c.Lat, c.Lon, k.Key)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Parse temp error: %v", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		log.Printf("Decode temp error: %v", err)
	}

	fmt.Printf("\nWeather in city: %f", w.Main.Temp)
	return w.Main.Temp
}
