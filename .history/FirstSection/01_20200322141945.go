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
	for _, ptc := range rna {
		rn = append(rn, ptc)

	}
	fmt.Println(rn)
	d := string([]rune(rn))
	fmt.Println(d)
}

func toString(x string) string {
	s := string(x)
	return s
}
