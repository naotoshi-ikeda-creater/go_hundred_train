package main

import "fmt"

func main() {
	cards := deck{newCard()}
	cards = append(cards, "six of the space")

	cards.print()

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamond"
}
