package main

import "fmt"

type bot interface {
	getGreeting() string
}

type enBot struct {
}

type esBot struct {
}

func main() {
	eb := enBot{}
	es := esBot{}

	printGreeting(eb)
	printGreeting(es)
}

func (enBot) getGreeting() string {
	return "Hi"
}

func (esBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
