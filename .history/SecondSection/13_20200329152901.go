package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("ファイルの読み込みを開始します。")

	arr1, arr2 := []string{}, []string{}

	read := func(file string) []string {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", file, err)
		}
		defer f.Close()

		lines := []string{}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if serr := scanner.Err(); serr != nil {
			fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", file, err)
		}

		return lines
	}

	read("col1.txt")
	read("col2.txt")

	write := func( []string, data string)
}
