package main

import "fmt"

func main() {
	p := "パトカー"
	t := "タクシー"
	rp := []rune(p)
	rt := []rune(t)
	fmt.Println(rp, rt)
	size := len(rp)
	var answer []rune
	for i, v := range rp rt {
		answer = append(answer, rp[v], rt[v])
	}
	fmt.Println(string(answer))
}
