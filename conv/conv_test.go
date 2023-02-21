package conv

import (
	"math"
	"testing"
)

func withinTolerance(a, b, error float64) bool {

	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	if b == 0 {
		return difference < error
	}

	return (difference / math.Abs(b)) < error
}

func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}

/*  SECOND TEST = CELCIUS TO FAHRENHEIT */
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0.0, want: 32.0},
	}
	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f; got %12.2f", tc.want, got)
		}
	}
}
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 273.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 32},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 273.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 0},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}
