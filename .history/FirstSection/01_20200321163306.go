package main

import "fmt"

func main() {
	cards := []string{newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamond"
}
