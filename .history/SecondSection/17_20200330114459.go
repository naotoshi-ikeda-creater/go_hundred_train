package main

import "fmt"

func main() {
	words := []string{"asuka", "asumatu", "asumatu", "asuko", "asuofasuko", "butachan", "asubu-", "asuasu"}
	for i, v := range words {
		if i == i*(len(words)/4) {
			fmt.Println(v + ",")
		} else {
			fmt.Println(v)
		}

	}
}
