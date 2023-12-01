package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for sum := 0.0; sum <= 20.0; {
		switch money := rand.Intn(3); money {
		case 0:
			sum += 0.05
		case 1:
			sum += 0.10
		case 2:
			sum += 0.25
		}
		fmt.Printf("%-05.2f\n", sum)
	}
}
