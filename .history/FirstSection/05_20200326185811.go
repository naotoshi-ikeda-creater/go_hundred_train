package main

import (
	"fmt"
	"strings"
)

func main() {
	w := "I am an NLPer"
	sep := strings.Fields(w)
	size := len(sep)
	for i := 0; i < size-1; i++ {
		fmt.Printf("%v", sep[i])
	}
}
