package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	arr1, arr2 := []string{}, []string{}

	read := func(filename string) []string {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filename, err)
		}
		defer f.Close()

		lines := []string{}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if serr := scanner.Err(); serr != nil {
			fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filename, err)
		}
		return lines
	}

	for _, v := range read("../data/hightemp.txt") {
		tmp := strings.Fields(v)
		arr1 = append(arr1, tmp[0])
		arr2 = append(arr2, tmp[1])
	}

	write := func(filename string, data []string) {
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		w := bufio.NewWriter(f)
		for _, v := range data {
			fmt.Fprint(w, v, "\n")
		}
		w.Flush()
	}

	write("col1.txt", arr1)
	write("col2.txt", arr2)

}
