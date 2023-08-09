package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com/",
		"http://youtube.com/",
		"http://github.com/",
		"http://google.com/"}

	c := make(chan string)

	for _, link := range links {
		go connect(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			connect(link, c)
		}(l)
	}

}

func connect(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}

	fmt.Println(link, "is connecting")
	c <- link
}
