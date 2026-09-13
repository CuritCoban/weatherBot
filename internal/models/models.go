package models

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
	ID   int
	City string
	Temp string
}
