package main

import "fmt"

func main() {
	p := "パトカー"
	t := "タクシー"
	rp := []rune(p)
	rt := []rune(t)
	fmt.Println(rp, rt)
	size := len(p)
}
