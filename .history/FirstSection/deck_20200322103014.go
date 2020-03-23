package main

import "fmt"

//create a new type of 'deck'
//which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"spades","diamonds","hearts","clubs"}
	cardValues := []string {"Ace","two", "three", "four"}

	for i, suit := range cardSuits {
		for j,value := range cardValues {
			cards := append(cards, suit+ "of"+value)
		}
			
		}
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
