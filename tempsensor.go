package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin int64
type sensor func() kelvin

func realSensor() kelvin {
	return kelvin(0)
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(100))
}

func calibrate(offset kelvin, s sensor) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func measureTempture(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Println(k, "Kelvin")
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	sensor := fakeSensor
	fmt.Println("Fake sensor:", sensor())
	sensor = calibrate(1, realSensor)
	fmt.Println("Real sensor:", sensor())
	func() {
		fmt.Println("SENSORS!!!")
	}()
}
