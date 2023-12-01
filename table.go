package main

import (
	"fmt"
)

type drawer func()
type tempconv func(c celsius) fahrenheit
type fahrenheit float64
type celsius float64

func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit((c * (9.0 / 5.0)) + 32.0)
}

func drawEqualLine(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
}

func drawTable(c tempconv) {
	drawEqualLine(21)
	fmt.Printf("|%8v|%10v|\n", "°C", "°F")
	drawEqualLine(21)
	for i := celsius(-40); i <= celsius(100); i += celsius(5) {
		fmt.Printf("|%8.2f|%10.2f|\n", i, c(i))
	}
	drawEqualLine(21)
}

func main() {
	drawTable(celsiusToFahrenheit)
}
