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
	for i := 0; i < size; i++ {
		answer = append(answer, rp[i], rt[i])
	}
}
