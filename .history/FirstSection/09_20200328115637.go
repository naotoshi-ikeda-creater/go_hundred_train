package main

import (
	"fmt"
	"strings"
)

// func shuffle(s string) string {
// 	rn := []rune(s)
// 	l := len(rn) - 1
// 	for i := 0; i < l; i++ {
// 		n := rand.Intn(l - i)
// 		rn[l-i], rn[n] = rn[n], rn[l-i]

// 	}
// 	return string(rn)

// }

// func expression(text []string) []string {
// 	var answer string

// 	for _, v := range text {
// 		if len(v) < 4 {
// 			answer = v + ""
// 		} else {
// 			rn := []rune(v)
// 			v = string(rn[0])
// 			s := string(rn[1 : len(rn)-1])
// 			v += shuffle(s)
// 			v += string(rn[len(rn)-1])
// 			answer = v + ""

// 		}
// 	}

// }

func main() {
	words := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	repl := strings.Split(words, "")
	fmt.Println(repl)

}
