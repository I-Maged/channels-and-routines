package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://facebook.com/",
		"http://youtube.com/",
		"http://github.com/",
		"http://google.com/",
		"http://sodfughoidsfuhgsidufhg.com/"}

	c := make(chan string)

	for _, link := range links {
		go connect(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func connect(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		c <- "Channel says link is down"
		return
	}

	fmt.Println(link, "is connecting")
	c <- "Channel says link is connecting"
}
