package main

import (
	"fmt"
	"strings"
)

func main() {
	w := "I am an NLPer"
	sep := strings.Fields(w)
	size := len(sep)
	for i := 0; i < size-1; i++ {
		fmt.Print(sep[i], sep[i+1], " ")
	}

	rn := []rune(w)
	fmt.Println(rn)
	sizee := len(rn)
	for i := 0, i<sizee - 1;i++ {
		fmt.Print(rn[i],rn[i+1],"")
	}
}
