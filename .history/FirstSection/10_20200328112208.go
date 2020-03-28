package main

import "fmt"

func main() {
	w := "work"
	rn := []rune(w)
	s := string(rn[1 : len(rn)-1])
	fmt.Println(s)
}
