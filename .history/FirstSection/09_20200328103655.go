package main

import (
	"fmt"
	"strings"
)

func shuffle(data []rune) []rune {

}

func main() {
	words := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	re := strings.Replace(words, ".", "", -1)
	rep := strings.Replace(re, ":", "", -1)
	repl := strings.Fields(rep)
	fmt.Println(repl)
}
