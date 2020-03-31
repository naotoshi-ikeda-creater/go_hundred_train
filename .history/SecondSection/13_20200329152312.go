package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("ファイルの読み込みを開始します。")

	read := func(file string) []string {
		f, err := os.Open(filename)
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
			fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filename, err)
		}

		return lines
	}
}
