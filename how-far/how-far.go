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
		cm, err := distanceOMatic.Distance_cm()
		if err != nil {
			panic(err)
		}
		fmt.Println("Cm:", cm)
	}
	fmt.Println("Woah nelly!")
}
