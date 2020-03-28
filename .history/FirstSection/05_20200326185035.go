package main

import (
	"fmt"
	"strings"
)

func main() {
	w := "I am an NLPer"
	sep := strings.Fields(w)
	size := len(sep)
	for i := 0; i < size; i++ {
		fmt.Printf("%v %v", sep[i], sep[i+1])
		break
	}
}
