package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gabrielcavalcantisiqueira/weather-app-go/internal/service"
)

type ViaCEPService struct{}

func NewViaCEPService() *ViaCEPService {
	return &ViaCEPService{}
}

func (s ViaCEPService) GetViaCEPInfo(cep string) (*service.ViaCEPResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data service.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	if hasError := strings.ToLower(data.Erro) == "true"; hasError {
		return nil, fmt.Errorf("not found")
	}
	return &data, nil
}
