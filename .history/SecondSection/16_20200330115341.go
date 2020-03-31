package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("ファイルの読み取りを開始します。")

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

	data = read("../data/hightemp.txt")
	// fmt.Println(data)

	n, _ := strconv.Atoi(os.Args[1])
	splitcnt := len(data) / n

	fmt.Printlnprintln(splitcnt)

	// fmt.Println(data[len(data)-*i:])

}
