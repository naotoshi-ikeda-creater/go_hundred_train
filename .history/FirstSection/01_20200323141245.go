package main

import "fmt"

func main() {
	s := []rune("パタトクカシー")

	var c []rune

	for i := len(s) - 1; i <= 6; {
		c = append(c, s[i])
	}

	fmt.Println(c)
}
