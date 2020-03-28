package main

import "fmt"

func main() {
	s := []rune("パタトクカシー")

	var c []rune

	fmt.Printf("%v", s)

	for i := len(s) - 7; i <= 6; i = i + 2 {
		c = append(c, s[i])
	}

	fmt.Printf("%v", string(c))
}
