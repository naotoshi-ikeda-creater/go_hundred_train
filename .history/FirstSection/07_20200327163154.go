package main

import (
	"fmt"
	"strings"
)

// 重複要素を取り除く
func RemoveDuplicate(xs []string) []string {
	ys := make([]string, 0, len(xs))
	for _, x := range xs {
		if !members(x, ys) {
			ys = append(ys, x)
		}
	}
	return ys
}

//与えられたslice(単語 or 文字)からngramを返す
func MakeNgrams(objs []string, n int, c string) []string {
	Ngrams := []string{}
	for i := 0; i < (len(objs) - n + 1); i++ {
		//fmt.Println((objs[i : i+n]))
		Ngrams = append(Ngrams, strings.Join(objs[i:i+n], c))
	}
	return RemoveDuplicate(Ngrams)
}
func main() {
	sent1 := "paraparaparadise"
	sent2 := "paragraph"
	fmt.Printf("%v\n", sent1)
	fmt.Printf("%v\n", sent2)

	X := MakeNgrams(strings.Split(sent1, ""), 2, "") //文字bigramの作成
	Y := MakeNgrams(strings.Split(sent2, ""), 2, "")
	fmt.Printf("X : %v\n", X)
	fmt.Printf("Y : %v\n", Y)

	//'se'がXに含まれるか
	if members("se", X) {
		fmt.Printf("'se' is member of X\n")
	} else {
		fmt.Printf("'se' isn't member of X\n")
	}

	//'se'がYに含まれるか
	if members("se", Y) {
		fmt.Printf("'se' is member of Y\n")
	} else {
		fmt.Printf("'se' isn't member of Y\n")
	}
}
