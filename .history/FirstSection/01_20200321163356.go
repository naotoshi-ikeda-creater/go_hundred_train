package main

import "fmt"

func main() {
	cards := []string{newCard()}
	cards = append(cards, "six of the space")
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamond"
}
