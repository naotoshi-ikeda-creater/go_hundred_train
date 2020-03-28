package main

import (
	"fmt"
	"strings"
)

func makeNgram(objs []string, n int, c string) []string {
	Ngrams := []string{}
	for i := 0; 1 < (len(objs) - n + 1); i++ {
		fmt.Println(objs[i : i+n])
		Ngrams = append(Ngrams, strings.Join(objs[i:i+n], c))
	}
}

func main() {

}
