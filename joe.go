package main

import (
	"fmt"
)

type speaker interface {
	speak() string
}

func speak(s speaker) {
	fmt.Println(s.speak())
}

type man struct {
	word string
}

func (m man) speak() string {
	return m.word
}

func main() {
	var Joe = man{word: "show you my hand"}
	speak(Joe)
}
