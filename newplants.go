package main

import (
	"fmt"
	"strings"
)

type Plants []string

func (pl Plants) AddNew() {
	for i, _ := range pl {
		pl[i] = strings.Join([]string{"New ", pl[i]}, "")
	}
}

func main() {

	var plants Plants = []string{
		"Mars",
		"Uranus",
		"Neptune",
	}

	fmt.Println(plants)
	plants.AddNew()
	fmt.Println(plants)

	fmt.Println("cap", cap(plants), "len", len(plants))
	plants = append(plants, "Mercury")
	fmt.Println("cap", cap(plants), "len", len(plants))
	makeplants := make([]string, 2, 3)
	fmt.Println(makeplants)
	fmt.Println("cap", cap(makeplants), "len", len(makeplants))
}
