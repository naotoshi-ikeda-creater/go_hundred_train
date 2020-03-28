package main

import (
	"fmt"
	"strings"
)

func main() {
	w := "I am an NLPer"
	sep := strings.Fields(w)
	size := len(sep)
	// for i := 0; i < size-1; i++ {
	// 	fmt.Print(sep[i], sep[i+1], " ")
	// }

	ws := strings.TrimSpace(w)
	fmt.Println(ws)
	rn := []rune(ws)
	fmt.Println(rn)
	sizee := len(rn)
	var answer []rune
	for i := 0; i < sizee-1; i++ {
		answer = append(answer, rn[i], rn[i+1])
	}

	fmt.Printf("%v", string(answer))
}
