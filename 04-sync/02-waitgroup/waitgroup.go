package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"https://epam.com",
		"https://amazon.com",
		"https://google.com",
	}

	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s - HTTP Status %d\n", url, res.StatusCode)
	}
}
