package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 10; i > 0; i-- {
		var year = 2000 + rand.Intn(100)
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			fmt.Println(year, "is leap year")
		} else {
			fmt.Println(year, "is not leap year!")
		}
	}
}
