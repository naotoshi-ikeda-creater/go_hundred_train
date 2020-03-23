package main

import (
	"fmt"
)

func main() {
	a := "パトカー"
	b := "タクシー"
	rna := []rune(a)
	rnb := []rune(b)
	fmt.Println(rna, rnb)
	var rn []rune
	for i := 0; i < size; i++ {
		rn = append(rn, rna[i], rnb[i])
	}
}

func toString(x string) string {
	s := string(x)
	return s
}
