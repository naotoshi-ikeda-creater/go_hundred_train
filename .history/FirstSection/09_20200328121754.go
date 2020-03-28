package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func shuffle(s string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rn := []rune(s)
	l := len(rn) - 1

	for i := 0; i < l; i++ {
		n := r.Intn(l - i)
		rn[l-i], rn[n] = rn[n], rn[l-i]

	}
	return string(rn)

}

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
	var words string:= "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	// repl := strings.Fields(words)
	// fmt.Println(repl)
	var answer string

	for _, v := range strings.Split(words, "") {
		if len(v) > 4 {
			rn := []rune(v)
			v = string(rn[0])
			s := string(rn[1 : len(rn)-1])
			v += shuffle(s)
			v += string(rn[len(rn)-1])

		}
		answer = v + ""
	}

	fmt.Println(answer)

}
