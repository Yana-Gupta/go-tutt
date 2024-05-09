package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://example.com/",
		"https://instagram.com",
		"https://facebook.com",
		"https://fac9e099book.com",
	}

	c := make(chan string)

	for _, link := range links {
		// These all will be the child process
		// Not given as much priority as the main Go Routine
		go checkLink(link, c)
	}

	for l := range c {

		go func(link string) {
			time.Sleep(4 * time.Second)
			go checkLink(link, c)
		}(l)
	}

	// The main routine doesn't care even if child routines are not done yet

	// Repeating Routines
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down 🔴")
		c <- link
		return
	}

	fmt.Println(link, "is up 🟢")
	c <- link
}

// Function Literal - Ananymous Function
