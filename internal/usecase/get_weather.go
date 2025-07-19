package usecase

import (
	"errors"
	"regexp"

	"github.com/gabrielcavalcantisiqueira/weather-app-go/internal/domain/entity"
	"github.com/gabrielcavalcantisiqueira/weather-app-go/internal/service"
)

var (
	ErrInvalidZipCode    = errors.New("invalid zipcode")
	ErrCannotFindZipCode = errors.New("can not find zipcode")
)

type GetWeatherUseCase struct {
	WeatherService service.WeatherServiceInterface
	ViaCEPService  service.ViaCEPServiceInterface
}

func NewGetWeatherUseCase(weatherService service.WeatherServiceInterface, viaCEPService service.ViaCEPServiceInterface) *GetWeatherUseCase {
	return &GetWeatherUseCase{WeatherService: weatherService, ViaCEPService: viaCEPService}
}

func (uc *GetWeatherUseCase) Execute(cep string) (*entity.Weather, error) {
	valid, _ := regexp.MatchString(`^\d{8}$`, cep)
	if !valid {
		return nil, ErrCannotFindZipCode
	}
	cepInfo, err := uc.ViaCEPService.GetViaCEPInfo(cep)
	if err != nil {
		if err.Error() == "not found" {
			return nil, ErrCannotFindZipCode
		}
		return nil, err
	}
	weather, err := uc.WeatherService.GetWeatherByCEP(cepInfo.Localidade)
	if err != nil {
		if err.Error() == "not found" {
			return nil, ErrCannotFindZipCode
		}
		return nil, err
	}

	return weather, nil
}
