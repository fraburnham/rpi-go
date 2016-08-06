package main

import (
	"fmt"
	"gpio"
	"gpio/sysfs"
)

const (
	inputPin = 6
)

func interruptWatch(eventCh chan gpio.InterruptEvent, ctrlCh chan bool) {
	for true {
		select{
		case event := <-eventCh:
			fmt.Println("Got event!")
			fmt.Println(event)
		case <-ctrlCh:
			return
		}
	}
}

func main() {
	eventCh := make(chan gpio.InterruptEvent, 2)
	ctrlCh := make(chan bool)
	pin, err := sysfs.NewSysfsInput(inputPin)
	if err != nil {
		panic(err)
	}

	err = pin.Interrupt("both", eventCh)
	if err != nil {
		panic(err)
	}

	go interruptWatch(eventCh, ctrlCh)

	// I think I recently saw a way to do this with runtime.???
	var junk string
	fmt.Scanln(&junk)
}
