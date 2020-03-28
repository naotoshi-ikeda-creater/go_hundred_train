package main

import (
	"fmt"
	"strings"
)

//nが含まれるかどうかを見極める関数
func include(n string, group []string) bool {
	for _, v := range group {
		if n == v {
			return true
		}
	}
	return false
}

func RemoveDuplicates(xa []string) []string {
	ys := make([]string, 0, len(xa))
	for _, x := range xa {
		if !include(x, ys) {
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
	return RemoveDuplicates(Ngrams)
}

//和集合
func makeUnion(set1, set2 []string) []string {
	set := []string{}
	set = append(set1, set2...) //sliceの結合
	return RemoveDuplicates(set)
}

//積集合
func MakeIntersection(set1, set2 []string) []string {
	set := []string{}
	//set1で取り出された文字列がset2の配列に含まれるかどうか
	for _, c := range set1 {
		if include(c, set2) {
			set = append(set, c)
		}
	}
	return set
}

//差集合
func MakeDifference(set1, set2 []string) []string {
	set := []string{}
	for _, v := range set1 {
		if !include(v, set2) {
			set = append(set, v)
		}
	}
	return set
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

	Union := makeUnion(X, Y)
	fmt.Printf("Union(X+Y): %v\n ", Union)

	Intersection := MakeIntersection(X, Y)
	fmt.Printf("Intersection(X,Y) : %v\n", Intersection)

	Difference := MakeDifference(X, Y)
	fmt.Printf("Difference(X,Y) : %v\n", Difference)

}
