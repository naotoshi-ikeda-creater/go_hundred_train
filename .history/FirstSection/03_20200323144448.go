package main

import "regexp"

func main() {
	text := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	new := regexp.MustCompile(`[A-Za-z]`)
}
