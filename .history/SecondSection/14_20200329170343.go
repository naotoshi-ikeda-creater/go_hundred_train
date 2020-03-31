package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("args count: %d\n", len(os.Args))
	fmt.Printf("args : %#v\n", os.Args)

	for i, v := range os.Args {
		fmt.Printf("args[%d] -> %s\n", i, v)
	}
}
