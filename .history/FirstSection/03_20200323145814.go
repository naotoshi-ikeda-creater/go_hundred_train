package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	re := strings.Replace(text, ",", "", -1)
	rep := strings.Replace(re, ".", "", -1)
	fmt.Println(rep)

}
