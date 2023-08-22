package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting

	return "Hi, there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting

	return "Hola!"
}

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) string {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) string {
// 	fmt.Println(sb.getGreeting())
// }
