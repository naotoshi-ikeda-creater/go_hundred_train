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
		for _, txy := range rnb {
			rn = append(rn, ptc, txy)
		}

	}
	fmt.Println(rn)
	d := string(rn)
}

func toString(x string) string {
	s := string(x)
	return s
}
