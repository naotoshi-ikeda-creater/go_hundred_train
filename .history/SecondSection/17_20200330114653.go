package main

import "fmt"

func main() {
	words := []string{"asuka", "asumatu", "asumatu", "asuko", "asuofasuko", "butachan", "asubu-", "asuasu"}
	mau := len(words) / 4
	for i, v := range words {
		if i == i*mau {
			fmt.Println(v + ",")
		} else {
			fmt.Println(v)
		}

	}
}
