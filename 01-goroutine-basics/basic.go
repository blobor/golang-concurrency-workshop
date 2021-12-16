package main

import (
	"fmt"
)

func printStr(str string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("\"%s\": %d\n", str, i)
	}
}

func main() {
	// spin new goroutine
	go func() {
		printStr("print from goroutine-1")
	}()

	// print in sync
	printStr("print sync")

	// one more goroutine
	func() {
		printStr("print from goroutine-2")
	}()
}
