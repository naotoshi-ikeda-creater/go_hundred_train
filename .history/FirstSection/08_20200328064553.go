package main

import (
	"fmt"
	"strings"
)

func cipher(input string) {
	rn := []rune(input)
	answer := []string{}

	for _, v := range rn {
		if v >= []rune("a")[0] && v <= []rune("z")[0] {
			c := 219 - v
			answer = append(answer, string(c))
		} else {
			answer = append(answer, string(v))
		}
	}
	result := strings.Join(answer, "")
}

func main() {
	message := "I really think that is true!"
	fmt.Println(cipher(message))
}
