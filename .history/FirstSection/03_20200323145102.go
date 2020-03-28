package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	new := regexp.MustCompile(`[A-Za-z]*$`)
	text = new.ReplaceAllString(text, "")
	fmt.Println(new)
}
