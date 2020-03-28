package main

import "fmt"

func cipher(input string) {
	rn := []rune(input)
	answer := []string{}

	for _, v := range rn {
		if v >= []rune("a")[0] && v <= []rune("z")[0] {
			c := 219 - v
			answer = append(answer, string(c))
		} else {
			answer = append(answer, v)
		}
	}
	return answer
}

func main() {
	message := "I really think that is true!"
	fmt.Println(cipher(message))
}
