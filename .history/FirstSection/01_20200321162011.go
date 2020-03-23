package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println("Hello there!")
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamond"
}
