package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"Haruka", "Chihaya", "Miki", "Yayoi", "Iori"}
	fmt.Println(a)
	r := strings.Join(a, ",")
	m := strings.Join(a, _)
	fmt.Println(r, m)

	rn := []rune(r)
	fmt.Println(rn)

}

func toString(x string) string {
	s := string(x)
	return s
}
