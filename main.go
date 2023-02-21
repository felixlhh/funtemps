package main

import (
	"flag"
	"fmt"

	"github.com/felixlhh/funtemps/conv"
)

var (
	fahr float64
	celc float64
	kelv float64
	in   string
	out  string
)

func main() {

	flag.Float64Var(&fahr, "F", 0, "Enter a temperature in Fahrenheit")
	flag.Float64Var(&celc, "C", 0, "Enter a temperature in Celsius")
	flag.Float64Var(&kelv, "K", 0, "Enter a temperature in Kelvin")
	flag.StringVar(&in, "in", "", "Enter the input temperature unit (F, C, K)")
	flag.StringVar(&out, "out", "", "Enter the desired output temperature unit (F, C, K)")
	flag.Parse()

	if in == "F" {
		switch out {
		case "C":
			result := conv.FahrenheitToCelsius(fahr)
			fmt.Printf("%.2f°F is equal to %.2f°C\n", fahr, result)
		case "K":
			result := conv.FahrenheitToKelvin(fahr)
			fmt.Printf("%.2f°F is equal to %.2fK\n", fahr, result)
		default:
			fmt.Println("Please enter a valid output temperature unit.")
		}
	} else if in == "C" {
		switch out {
		case "F":
			result := conv.CelsiusToFahrenheit(celc)
			fmt.Printf("%.2f°C is equal to %.2f°F\n", celc, result)
		case "K":
			result := conv.CelsiusToKelvin(celc)
			fmt.Printf("%.2f°C is equal to %.2fK\n", celc, result)
		default:
			fmt.Println("Please enter a valid output temperature unit.")
		}
	} else if in == "K" {
		switch out {
		case "C":
			result := conv.KelvinToCelsius(kelv)
			fmt.Printf("%.2fK is equal to %.2f°C\n", kelv, result)
		case "F":
			result := conv.KelvinToFahrenheit(kelv)
			fmt.Printf("%.2fK is equal to %.2f°F\n", kelv, result)
		default:
			fmt.Println("Please enter a valid output temperature unit.")
		}
	} else {
		fmt.Println("Please enter a valid input temperature unit.")
	}
}
