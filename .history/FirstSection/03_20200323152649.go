package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	re := strings.Replace(text, ",", "", -1)
	rep := strings.Replace(re, ".", "", -1)
	fmt.Println(rep)
	answer := strings.Fields(rep)
	fmt.Println(answer, len(answer))
	size := len(answer)
	for i := 0; i < size; i++ {
		fmt.Println(answer, "%v" utf8.RuneCountInString(answer))
	}
}
