package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width        = 100
	height       = 40
	seed_percent = 2
	turn         = 25*10 //50 fps * 10 seconds
	sleep_ms     = 40
)

type Universe [][]bool

// Method of type Universe: Show
// print slice to terminal
func (u Universe) Show() {
	fmt.Printf("\n")
	for i := 0; i < width*2+2; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
	for i := range u {
		fmt.Printf("|")
		for _, value := range u[i] {
			if value { //if cell is alive
				fmt.Printf("* ")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("|\n")
	}
	for i := 0; i < width*2+2; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

// Method of type Universe: Seed
// Randomly change some cell to true based on give seed percentage
func (u Universe) Seed(s int) {
	for i := range u {
		for ii, _ := range u[i] {
			if rand.Intn(100) < s { //if cell is alive
				u[i][ii] = true
			} else {
				u[i][ii] = false
			}
		}
	}
}

// Method of type Universe: Alive
// check if a sell in (x,y) stays alive
func (u Universe) Alive(x, y int) bool {
	// i := y
	// ii := x
	count := 0

	if u[y][(x+1+width)%width] {
		count++
	}
	if u[(y+1+height)%height][(x+1+width)%width] {
		count++
	}
	if u[(y-1+height)%height][(x+1+width)%width] {
		count++
	}
	if u[y][(x-1+width)%width] {
		count++
	}
	if u[(y+1+height)%height][(x-1+width)%width] {
		count++
	}
	if u[(y-1+height)%height][(x-1+width)%width] {
		count++
	}
	if u[(y+1+height)%height][x] {
		count++
	}
	if u[(y-1+height)%height][x] {
		count++
	}

	val := false
	switch count {
	case 0, 1, 4, 5, 6, 7, 8:
		val = false
	case 2, 3:
		val = true
	default:
		print("ERROR in Alive()")
	}
	return val
}

func (u Universe) Next() Universe {
	nextu := NewUniverse()
	for i := range u {
		for ii, _ := range u[i] {
			if u.Alive(ii, i) { //if cell is alive
				nextu[i][ii] = true
			} else {
				nextu[i][ii] = false
			}
		}
	}
	return nextu
}

func NewUniverse() Universe {
	sliceU := make([][]bool, height)
	for i := range sliceU {
		sliceU[i] = make([]bool, width)
	}
	return Universe(sliceU)
}

func main() {
	universe := NewUniverse()
	universe.Seed(seed_percent)
	for t := 0; t < turn; t++ {
		universe = universe.Next()
		universe.Show()
		time.Sleep(time.Millisecond * sleep_ms)
	}
}
