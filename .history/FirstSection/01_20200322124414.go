package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"Haruka", "Chihaya", "Miki", "Yayoi", "Iori"}
	fmt.Println(a)
	r := strings.Join(a, ",")
	fmt.Println(r)

}

func toString(x string) string {
	s := string(x)
	return s
}
