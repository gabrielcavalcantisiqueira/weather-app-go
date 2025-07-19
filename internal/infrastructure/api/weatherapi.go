package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gabrielcavalcantisiqueira/weather-app-go/internal/domain/entity"
)

type WeatherService struct {
	APIKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		APIKey: apiKey,
	}
}

func (s *WeatherService) GetWeatherByCEP(location string) (*entity.Weather, error) {
	escapedLocation := url.QueryEscape(location)
	apiURL := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", s.APIKey, escapedLocation)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var data struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	tempC := data.Current.TempC
	weather := entity.NewWeather(tempC)
	return weather, nil
}
