package main

import (
	"rangefinder"
	"fmt"
)

const (
	triggerPin = 5
	signalPin = 6
)

func main() {
	// create the range finder and run a loop reporting the distance to an object
	distanceOMatic, err := rangefinder.NewHCSRO4(triggerPin, signalPin)
	if err != nil {
		panic(err)
	}
	for true {
		distanceOMatic.Distance_cm()
	}
	fmt.Println("Woah nelly!")
}
