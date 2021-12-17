package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	go calculateRand(ch)

	val := <-ch
	fmt.Println(val)
}

func calculateRand(ch chan int) {
	fmt.Println("Starting hard working process")

	time.Sleep(10 * time.Second)

	ch <- rand.Intn(1 << 31)
}
