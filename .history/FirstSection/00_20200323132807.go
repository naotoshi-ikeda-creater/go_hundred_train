package main

import "fmt"

func main() {
	s := []rune("stressed")
	fmt.Println(s)
	l := len(s)
	fmt.Println(l)
	// var c []rune
	var i int = l - 1
	var answer []rune
	fmt.Println(i)
	for i := len(s) - 1; i >= 0; i-- {
		answer = append(s[i]rune)
	}

}
