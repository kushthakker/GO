package main

import "fmt"

type bot interface {
	getGreeting(int) string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb, 1)
	printGreeting(sb, 2)
}

func printGreeting(b bot, i int){
	fmt.Println(b.getGreeting(i))
}

func (englishBot) getGreeting(i int) string {
	return fmt.Sprintf("Hi there %d", i)
}
func (spanishBot) getGreeting(i int) string {
	return "Hola!"
}
