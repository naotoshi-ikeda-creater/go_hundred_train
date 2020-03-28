package main

import "fmt"

func template(x int, y string, z int) {
	fmt.Printf("%v時の%vは%v", x, y, z)

}

func main() {
	template(12, "気温", 22.4)
}
