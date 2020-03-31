package main

import (
	"bufio"
	"fmt"
	"os"
)

func include(word string, uni []string) bool {
	for _, v := range uni {
		if word == v {
			return true
		}

	}
	return false
}

func main() {

	data := []string{}

	read := func(path string) []string {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", path, err)
		}
		defer f.Close()

		lines := []string{}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if serr := scanner.Err(); serr != nil {
			fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", path, err)
		}
		return lines
	}

	data = read("col1.txt")

	arr := []string{}

	for _, v := range data {
		if !include(v, arr) {
			fmt.Println(v)
		}
	}

}
