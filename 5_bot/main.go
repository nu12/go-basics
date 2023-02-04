package main

import "fmt"

type bot interface {
	// Every type that implements the getGreeting function
	// returning a string is also part of the bot interface
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	printGreeting(englishBot{})
	printGreeting(spanishBot{})

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
