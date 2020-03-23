package main 


import

func main() {
	
	a := "パトカー"
	b := "タクシー"
	rna := []rune(a)
	rnb := []rune(b)
	var rn []rune
	for i, _ := range rna, rnb{
		rn = append(rn, rna[i], rnb[i])
	}
	fmt.Printf("%s\n", string(rn))
}