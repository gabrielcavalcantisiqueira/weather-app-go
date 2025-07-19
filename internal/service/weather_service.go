package service

import "github.com/gabrielcavalcantisiqueira/weather-app-go/internal/domain/entity"

type WeatherServiceInterface interface {
	GetWeatherByCEP(cep string) (*entity.Weather, error)
}
