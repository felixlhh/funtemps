package conv

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(c float64) float64 {
	return c*(9.0/5.0) + 32.0
}

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32.0) * 5.0 / 9.0
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

// Konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(f float64) float64 {
	return (f-32)*(5.0/9.0) + 273.15
}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(k float64) float64 {
	return (k-273.15)*(9.0/5.0) + 32.0
}
