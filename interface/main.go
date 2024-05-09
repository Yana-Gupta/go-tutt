package main

import "fmt"

type bot interface {
	// Any type on our module that implement generateGreeting is a part of this interface bot
	generateGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) generateGreeting() string {
	return "Hello there!"
}

func (spanishBot) generateGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.generateGreeting())
}
