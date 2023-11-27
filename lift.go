package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const chance_percent = 5

	var countdown = 10
	for countdown > 0 {
		fmt.Println(countdown)
		if rand.Intn(100) < chance_percent {
			print("BOOM!!!")
			break
		}
		countdown--
	}

	if countdown == 0 {
		print("GOOOO")
	}
}
