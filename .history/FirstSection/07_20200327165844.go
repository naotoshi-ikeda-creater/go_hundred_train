package main

import (
	"fmt"
	"strings"
)

// 重複要素を取り除く

//与えられたslice(単語 or 文字)からngramを返す
func MakeNgrams(objs []string, n int, c string) []string {
	Ngrams := []string{}
	for i := 0; i < (len(objs) - n + 1); i++ {
		//fmt.Println((objs[i : i+n]))
		Ngrams = append(Ngrams, strings.Join(objs[i:i+n], c))
	}
	return Ngrams
}

func makeUnion (set1, set2 []string) []string[
	set := []string {}
	set = append(set1,set2...)//sliceの結合
	return set
]
func main() {
	sent1 := "paraparaparadise"
	sent2 := "paragraph"
	fmt.Printf("%v\n", sent1)
	fmt.Printf("%v\n", sent2)

	X := MakeNgrams(strings.Split(sent1, ""), 2, "") //文字bigramの作成
	Y := MakeNgrams(strings.Split(sent2, ""), 2, "")
	fmt.Printf("X : %v\n", X)
	fmt.Printf("Y : %v\n", Y)

	Union := makeUnion(X,Y)
	fmt.Printf("Union(X+Y): %v\n ", Union)

}
