package main

import "fmt"

func main() {
	cards := deck{newCard()}
	cards = append(cards, "six of the space")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamond"
}
