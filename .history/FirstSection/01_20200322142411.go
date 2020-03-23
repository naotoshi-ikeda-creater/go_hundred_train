package main

import (
	"fmt"
)

func main() {
	a := "パトカー"
	b := "タクシー"
	rna := []rune(a)
	rnb := []rune(b)
	size := len(rna)
	var rn []rune
	for i := 0; i < size; i++ {
		rn = append(rn, rna[i], rnb[i])
	}
	fmt.Printf("%s\n", string(rn))
}

func toString(x string) string {
	s := string(x)
	return s
}
