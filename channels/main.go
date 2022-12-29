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
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(10 * time.Second)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "might be down")
		c <- l
		return
	}
	fmt.Println(l, "is up")
	c <- l
}
