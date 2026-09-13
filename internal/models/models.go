package models

import "time"

type CoordinateSlice []struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Coordinate struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

type WeatherDB struct {
	ChatID    int64     `gorm:"chat_id"`
	City      string    `gorm:"city"`
	Temp      float64   `gorm:"temp"`
	Lon       float64   `json:"lon"`
	Lat       float64   `json:"lat"`
	CreatedAt time.Time `gorm:"created_at"`
}
