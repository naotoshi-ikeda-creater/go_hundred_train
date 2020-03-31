package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	arr3 := []string{}

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
		arr3 = append(arr3, tmp[2])
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
	data1 := read("../data/hightemp.txt")

	//３コラム目のファイルを作成する
	write("col3.txt", arr3)

	//3コラム目のファイルを読み込む
	data3 := read("col3.txt")

}
