package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for sum := 0; sum <= 2000; {
		switch money := rand.Intn(3); money {
		case 0:
			sum += 5
		case 1:
			sum += 10
		case 2:
			sum += 25
		}
		fmt.Printf("$%v.%02v\n", sum/100, sum%100)
	}
}
