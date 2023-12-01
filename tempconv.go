package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type fahrenheit float64

func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit((c * (9.0 / 5.0)) + 32.0)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k - 273.15*(9.0/5.0)) + 32.0)
}

func main() {
	var kel kelvin = 0.0
	fmt.Println(kel.celsius())
	fmt.Println(kel.fahrenheit())
}
