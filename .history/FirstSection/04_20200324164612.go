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
	
	for i, v := range rep {
		for _, n := range number {
			if i+1 == n {
				// fmt.Println(v)
				fmt.Printf("%v\n", v[:1])
			} else  i+1 != n{
				fmt.Printf("%v\n", v[:2])
				fmt.Println(i)
			}

		}
	}
}
