package entity

type Weather struct {
	TempC float64
	TempF float64
	TempK float64
}

func NewWeather(tempC float64) *Weather {
	return &Weather{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273.15,
	}
}
