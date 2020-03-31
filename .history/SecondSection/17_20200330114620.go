package main

import "fmt"

func main() {
	mau := len(words) / 4
	words := []string{"asuka", "asumatu", "asumatu", "asuko", "asuofasuko", "butachan", "asubu-", "asuasu"}
	for i, v := range words {
		if i == i*mau {
			fmt.Println(v + ",")
		} else {
			fmt.Println(v)
		}

	}
}
