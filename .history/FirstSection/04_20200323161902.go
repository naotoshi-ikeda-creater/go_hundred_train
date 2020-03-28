package main

import "strings"

func main() {
	text := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	re := strings.Replace(text, ".", "", -1)
	rep := strings.Fields(re)
	var answer []string
	for _, v := range rep {
		if i == 1, 5, 6, 7, 8, 9, 15, 16, 19 {
			answer = append(string([]rune(v)[:2]))
		}else{
			answer = append(string([]rune(v)[0]))
		}
	}
}
