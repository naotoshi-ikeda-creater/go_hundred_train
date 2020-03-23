package main

func main() {
	cards := deck{newCard()}
	cards = append(cards, "six of the space")

	cards.print()

}

func newCard() string {
	return "Five of Diamond"
}
