package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	fmt.Println("Start counting")
	go countTo(10, ch)

	for v := range ch {
		fmt.Printf("%d\n", v)
	}

	fmt.Println("Stop counting")
}

func countTo(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}

	close(ch)
}
