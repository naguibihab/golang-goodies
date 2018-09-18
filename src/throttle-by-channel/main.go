package main

import (
	"log"
	"time"
)

// Make a channel in the length of our concurrency limit
const conLimit = 5
var doubleFunChannel = make(chan bool, conLimit)

func main() {

	for v := 0; v < 10; v++ {
		go func(v int) {
			doublev := callDouble(v)
			log.Printf("Thread %d returned: %d", v, doublev)
		}(v)
	}

	time.Sleep(time.Second * 10)
}

func callDouble(v int) int {
	// Add a value to the channel (it doesn't matter what value)
	doubleFunChannel <- true
	
	// Remove value from the channel when the function is done
	defer func() { <-doubleFunChannel  }()
	
	return double(v)
}

func double(v int) int {
	time.Sleep(time.Second)
	return v * 2
}
