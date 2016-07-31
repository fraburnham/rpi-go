package main

import (
	"gpio"
	"gpio/events"
	"fmt"
)

const (
	inputPin = 6
)

func triggerWatch(eventCh chan events.EdgeEvent, ctrlCh chan bool) {
	for true {
		event := <-eventCh
		fmt.Println("Got event!")
		fmt.Println(event)
	}
}

func main() {
	pin, err := gpio.NewRpiInput(inputPin)
	if err != nil {
		panic(err)
	}

	eventCh, eCtrlCh := events.StartEdgeTrigger(pin)
	tCtrlCh := make(chan bool)

	go triggerWatch(eventCh, tCtrlCh)

	var junk string
	fmt.Scanln(&junk)
	eCtrlCh <- true
	tCtrlCh <- true
}
