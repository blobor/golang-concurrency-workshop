package main

import (
	"fmt"
	"sync"
)

func main() {
	once := sync.Once{}

	for i := 0; i < 1000; i++ {
		once.Do(func() {
			fmt.Println("Run Once!")
		})
	}

	fmt.Println("Finish test")
}
