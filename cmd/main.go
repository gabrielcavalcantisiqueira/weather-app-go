package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gabrielcavalcantisiqueira/weather-app-go/internal/infrastructure/api"
	"github.com/gabrielcavalcantisiqueira/weather-app-go/internal/interface/handler"
	"github.com/gabrielcavalcantisiqueira/weather-app-go/internal/usecase"
)

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("API Key não definida")
		return
	}
	weatherService := api.NewWeatherService(apiKey)
	viaCEPService := api.NewViaCEPService()
	useCase := usecase.NewGetWeatherUseCase(weatherService, viaCEPService)
	h := handler.NewWeatherHandler(useCase)
	http.Handle("/weather", h)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
