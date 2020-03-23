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
}

func toString(x string) string {
	s := string(x)
	return s
}
