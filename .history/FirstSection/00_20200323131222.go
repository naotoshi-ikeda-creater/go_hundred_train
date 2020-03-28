package main

import "fmt"

func main() {
	s := []rune("stressed")
	fmt.Println(s)
	l := len(s)
	fmt.Println(l)
	var c []rune
	for i, n := range s {
		c = append(s[i-1])

	}

