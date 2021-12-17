package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	buffer := []byte{}

	fmt.Println("Start filling buffer")

	for i := 0; i < 1000; i++ {
		go func() {
			buffer = append(buffer, byte(rand.Int31n(1<<7)))
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Printf("buffer length = %d\n", len(buffer))
}
