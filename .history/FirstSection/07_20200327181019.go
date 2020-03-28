package main

import "fmt"

func template(x int, y string, z int) string {
	answer := fmt.Printf("%v時の%vは%v", x, y, z)
}
