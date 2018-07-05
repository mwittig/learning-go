package main

import (
	"fmt"
)

func main() {
	x := 4
	y := 16

	// Create a channel that consumes and outputs boolean values
	// Channels are used as synchronization points for go routines
	// Channels may also be used to return some result from a go routine
	operationDone := make(chan bool)

	go func() {
		x = x * y

		// push some boolean value to the channel
		operationDone <- true
	}()

	// wait for a message to be pushed into the channel
	// in this case we do not check the value returned by the channel
	<-operationDone

	y = x * x

	fmt.Printf("x = %d, y = %d\n", x, y)
}