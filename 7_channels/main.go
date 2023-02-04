package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"google.com",
		"golang.org",
		"amazon.com",
		"stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string, channel chan string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get("http://" + link)
	if err != nil {
		fmt.Println("Error:", err)
		c <- link
		return
	}
	fmt.Println(link, "ok")
	c <- link
}
