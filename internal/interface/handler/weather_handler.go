package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielcavalcantisiqueira/weather-app-go/internal/usecase"
)

type WeatherHandler struct {
	UC *usecase.GetWeatherUseCase
}

func NewWeatherHandler(uc *usecase.GetWeatherUseCase) *WeatherHandler {
	return &WeatherHandler{
		UC: uc,
	}
}

func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	weather, err := h.UC.Execute(cep)
	if err != nil {
		if err == usecase.ErrInvalidZipCode {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if err == usecase.ErrCannotFindZipCode {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]float64{
		"temp_C": weather.TempC,
		"temp_F": weather.TempF,
		"temp_K": weather.TempK,
	})
}
