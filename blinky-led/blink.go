package main

import (
	"fmt"
	"gpio"
	"gpio/sysfs"
	"time"
)

func timeoutGenerator(ctrlCh chan bool, timeoutCh chan bool, timeout time.Duration) {
	fmt.Println("Starting timeoutGenerator")
	for true {
		select {
		case <-ctrlCh:
			fmt.Println("timeoutGenerator dying")
			return
		default:
			time.Sleep(timeout)
			timeoutCh <- true
		}
	}
}

func blink(ctrlCh chan bool, timeoutCh chan bool, pin gpio.GPIO) {
	fmt.Println("Starting blink")
	offset := 0

	for true {
		select {
		case <-ctrlCh:
			fmt.Println("blink dying")
			return
		case <-timeoutCh:
			offset++
			err := pin.WriteValue(offset % 2)
			if err != nil {
				panic(err.Error())
			}
		}
	}

}

func main() {
	ctrlCh := make(chan bool) //oh snap, I bet channels need to be closed
	timeoutCh := make(chan bool)
	timeout := time.Duration(1 * time.Second)
	pin, err := sysfs.NewSysfsOutput(4)
	defer pin.Close()

	if err != nil {
		panic(err.Error())
	}

	go blink(ctrlCh, timeoutCh, pin)
	go timeoutGenerator(ctrlCh, timeoutCh, timeout)

	fmt.Println("Press enter to quit")
	junk := ""
	fmt.Scanln(&junk)
	ctrlCh <- true
}
