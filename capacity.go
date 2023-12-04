package main

import (
	"fmt"
)

func main() {
	clip := make([]string, 5)
	fmt.Println("cap", cap(clip), "len", len(clip))

	for i := 0; i < 100; i++ {
		clip = append(clip, "")
		fmt.Println("cap", cap(clip), "len", len(clip))
	}
}
