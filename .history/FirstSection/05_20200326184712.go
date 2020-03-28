package main

import (
	"fmt"
	"strings"
)

func main() {
	w := "I am an NLPer"
	sep := strings.Fields(w)

	for _, value := range sep {
		fmt.Printf("%v", value[i], value[i+1])
	}
}
