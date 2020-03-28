package main

import "fmt"

func main() {
	s := []rune("stressed")
	fmt.Println(s)
	l := len(s)
	fmt.Println(l)
	// var c []rune
	var i int = l - 1

	fmt.Println(i)
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(string(s[i]))
	}

}
