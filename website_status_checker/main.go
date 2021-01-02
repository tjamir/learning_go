package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for link := range c {
		go func(link string, channel chan string) {
			time.Sleep(10 * time.Second)
			go checkLink(link, channel)
		}(link, c)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error checking site status : %v\n", link)
	} else {
		fmt.Printf("Site seems to be ok: %v\n", link)
	}
	channel <- link
}
