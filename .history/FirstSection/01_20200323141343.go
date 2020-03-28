package main

import "fmt"

func main() {
	s := []rune("パタトクカシー")

	var c []rune

	fmt.Printf("%v s", s)

	for i := len(s) - 1; i <= 6; i++ {
		c = append(c, s[i])
	}

	fmt.Println(c)
}
