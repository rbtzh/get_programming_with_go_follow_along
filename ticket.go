package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const distance = 62100000 // km
	fmt.Printf("%-25v %-10v %-10v %-10v\n\n", "Company", "Duration", "Type", "Price")
	for i := 10; i > 0; i-- {
		var company string
		switch i := rand.Intn(3); i {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		var flightTypeString string
		var flightType = rand.Intn(2) + 1 // 1 or 2
		if flightType == 1 {
			flightTypeString = "Single"
		} else if flightType == 2 {
			flightTypeString = "Double"
		} else {
			fmt.Println("ERROR")
		}

		var speed = rand.Intn(30-16+1) + 16 // 16 ~ 30 km
		var price = (3600.0 + (float64(speed-16)/float64(30-16))*float64(5000-3600)) * float64(flightType)
		var duration = (distance / speed) / (60 * 60 * 24) //day
		fmt.Printf("%-25v %-10v %-10v %-10v\n", company, duration, flightTypeString, price)
	}
}
