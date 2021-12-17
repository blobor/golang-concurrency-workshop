package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	fmt.Println("Start computing random values")

	go getRandomValue1(ch1)
	go getRandomValue2(ch2)

	select {
	case v1 := <-ch1:
		fmt.Printf("got %d from channel 1\n", v1)
	case v2 := <-ch2:
		fmt.Printf("got %d from channel 2\n", v2)
	}

	// what will be an output?
}

func getRandomValue1(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 16
}

func getRandomValue2(ch chan int) {
	time.Sleep(5 * time.Second)
	ch <- 86
}
