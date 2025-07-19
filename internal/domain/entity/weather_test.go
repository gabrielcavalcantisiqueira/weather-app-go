package entity

import (
	"math"
	"testing"
)

func TestNewWeather(t *testing.T) {
	tempC := 25.0
	expectedF := 77.0
	expectedK := 298.15

	weather := NewWeather(tempC)

	if weather.TempC != tempC {
		t.Errorf("Expected TempC: %.2f, got: %.2f", tempC, weather.TempC)
	}

	if !almostEqual(weather.TempF, expectedF) {
		t.Errorf("Expected TempF: %.2f, got: %.2f", expectedF, weather.TempF)
	}

	if !almostEqual(weather.TempK, expectedK) {
		t.Errorf("Expected TempK: %.2f, got: %.2f", expectedK, weather.TempK)
	}
}

func almostEqual(a, b float64) bool {
	const epsilon = 0.00001
	return math.Abs(a-b) < epsilon
}
