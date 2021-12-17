package main

import (
	"fmt"
	"time"
)

var counter int32

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	time.Sleep(2 * time.Second)

	fmt.Printf("counter = %d\n", counter)

	// do we have a race here?
}

func increment() {
	counter += 1
}
