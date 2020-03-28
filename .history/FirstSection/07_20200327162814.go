package main

import (
	"strings"
)

// n がスライスに含まれているか
func members(n string, xs []string) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

// 重複要素を取り除く
func RemoveDuplicate(xs []string) []string {
	ys := make([]string, 0, len(xs))
	for _, x := range xs {
		if !member(x, ys) {
			ys = append(ys, x)
		}
	}
	return ys
}

//与えられたslice(単語 or 文字)からngramを返す
func MakeNgram(objs []string, n int, c string) []string {
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

	X := MakeNgram(strings.Split(sent1, ""), 2, "") //文字bigramの作成
	Y := MakeNgram(strings.Split(sent2, ""), 2, "")
	fmt.Printf("X : %v\n", X)
	fmt.Printf("Y : %v\n", Y)

	Union := MakeUnion(X, Y) //和集合
	fmt.Printf("Union(X,Y) : %v\n", Union)

	Intersection := MakeIntersection(X, Y) //積集合
	fmt.Printf("Intersection(X,Y) : %v\n", Intersection)

	Difference := MakeDifference(X, Y) //差集合
	fmt.Printf("Difference(X,Y) : %v\n", Difference)

	//'se'がXに含まれるか
	if member("se", X) {
		fmt.Printf("'se' is member of X\n")
	} else {
		fmt.Printf("'se' isn't member of X\n")
	}

	//'se'がYに含まれるか
	if member("se", Y) {
		fmt.Printf("'se' is member of Y\n")
	} else {
		fmt.Printf("'se' isn't member of Y\n")
	}
