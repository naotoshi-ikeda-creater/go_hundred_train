package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func shuffle(pro []rune) {
	for i := range pro {
		r := rand.Intn(i + 1)
		pro[i], pro[r] = pro[r], pro[i]
	}

}

func expression(text []string) []string {
	answer := []string{}

	for _, v := range text {
		if len(v) < 4 {
			answer = append(answer, string(v))
		} else {
			rn := []rune(v)
			s := string(rn[1 : len(rn)-1])

		}
	}

}

func main() {
	words := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	re := strings.Replace(words, ".", "", -1)
	rep := strings.Replace(re, ":", "", -1)
	repl := strings.Fields(rep)
	fmt.Println(repl)
}
