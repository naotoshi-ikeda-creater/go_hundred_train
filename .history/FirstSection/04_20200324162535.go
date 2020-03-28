package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	re := strings.Replace(text, ".", "", -1)
	rep := strings.Fields(re)
	// var answer []string
	number := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	for i, v := range rep {
		for _, n := range number {
			if i+1 == n {
				// fmt.Println(v)
				fmt.Printf("%v\n", v[:1])
				rn := []rune(v)
				fmt.Println(rn)

			}

		}
	}
}
