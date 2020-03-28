package main

import (
	"fmt"
	"strings"
)

func member(n string, xs []string) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

// 重複要素を取り除く
func RemoveDuplicates(xs []string) []string {¥
	ys := make([]string, 0, len(xs))
	for _, x := range xs {
		if !member(x, ys) {
			ys = append(ys, x)
		}
	}
	return ys
}

func makeNgram(objs []string, n int, c string) []string {
	Ngrams := []string{}
	for i := 0; 1 < (len(objs) - n + 1); i++ {
		fmt.Println(objs[i : i+n])
		Ngrams = append(Ngrams, strings.Join(objs[i:i+n], c))
	}
	return RemoveDuplicates(Ngrams)
}

func main() {
	sent1 := "paraparaparadise"

	X := MakeNgram(strings.Split(sent1, ""), 2, "")

	fmt.Printf("X : %v\n", X)
}
