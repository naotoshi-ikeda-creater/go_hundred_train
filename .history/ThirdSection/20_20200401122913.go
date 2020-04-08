package main

import "fmt"

func f(point *int) {
	*point = 100
}

func main() {
	var x int
	f(&x)
	fmt.Println(*x)

}
