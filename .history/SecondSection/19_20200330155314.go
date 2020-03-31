package main

import (
	"log"
	"sort"
)

// 独自の構造体
type Person struct {
	Name string
	Age  uint8
}

// 構造体のスライス
type People []Person

// 以下インタフェースを満たす

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func main() {
	var people People = []Person{
		{"John", 22},
		{"Tom", 20},
		{"Emily", 18},
	}

	// sort.Sort に渡すだけ
	sort.Sort(people)   // return はされない点には注意
	log.Println(people) // [{Emily 18} {Tom 22} {John 20}]
}
