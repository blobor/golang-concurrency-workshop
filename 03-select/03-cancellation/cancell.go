package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// if user hit CTRL-C, we will recieve notificaltion
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	fibCh := make(chan int)

	fmt.Println("Starting fibonacci numbers generation")
	go genFibonacci(fibCh)

	for {
		select {
		case v := <-fibCh:
			time.Sleep(1 * time.Second)
			fmt.Println(v)
		case <-signalCh:
			fmt.Println("Cancelled: shutdown CLI")
			os.Exit(1)
		}
	}

	// how to make a timeout?
}

func genFibonacci(ch chan int) {
	a, b := 0, 1
	for {
		ch <- a
		a, b = b, a+b
	}
}
