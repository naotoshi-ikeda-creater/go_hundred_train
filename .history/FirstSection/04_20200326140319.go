package main

import (
	"strings"
)

func main() {
	text := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	re := strings.Replace(text, ".", "", -1)
	rep := strings.Fields(re)

	// var answer []string
	number := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}

	m := make(map[string]int)
	for i, v := range rep {
		for _, n := range number {
			if i+1 == n {
				m[string(v[:1])] = i + 1
			} else {
				m[string(v[:2])] = i + 2

			}

		}
	}
}
