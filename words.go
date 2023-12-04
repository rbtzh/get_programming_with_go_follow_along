package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot!"

	words := strings.Split(strings.ReplaceAll(text, ",", ""), " ")
	wordmap := make(map[string]int)
	for _, v := range words {
		wordmap[v] += 1
	}

	for key, value := range wordmap {
		if value > 1 {
			fmt.Println(key, ":", value)
		}
	}
}
