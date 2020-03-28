package main

import "fmt"

func main() {
	s := []rune("パタトクカシー")

	var c []rune

	fmt.Printf("%v s", s)

	for i := len(s) - 6; i <= 6; i++ {
		c = append(c, s[i])
	}

	fmt.Printf("%v", string(c))
}
