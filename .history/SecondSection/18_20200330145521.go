package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

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

	data1 := read("../data/hightemp.txt")

	r := string[]{}

	for _, v := range data1 {
		data3 := strings.Fields(v)
		if r[2] >= v[2] {
			r = append(r,v)
		}else{
			continue
		}
	}

	fmt.Println(r)
}
