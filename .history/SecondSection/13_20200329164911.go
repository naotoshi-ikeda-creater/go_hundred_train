package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("ファイルの読み込みを開始します。")

	arr1, arr2 := []string{}, []string{}

	reading := func(file string) []string {
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

	arr1 = reading("col1.txt")
	arr2 = reading("col2.txt")

	writing := func(data1, data2 []string, text string) {
		f, err := os.Create(text)
		if err != nil {
			panic(err)
		}

		w := bufio.NewWriter(f)
		for i, _ := range data1 {
			fmt.Fprintf(w, data1[i], "\b", data2[i], "\n")
		}

		w.Flush()

	}

	writing(arr1, arr2, "marge.txt")
}
