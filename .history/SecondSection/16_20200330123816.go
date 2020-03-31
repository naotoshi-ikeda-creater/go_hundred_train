package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// This function is to convert slice of string to int
func convert(sli []string) ([]int, error) {
	si := make([]int, 0, len(sli))
	for _, a := range sli {
		i, err := strconv.Atoi(a)
		if err != nil {
			return sli, err
		}
		sli = append(sli, i)
	}
}

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
	if len(os.Args) < 2 {
		fmt.Println("ERROR: 引数を指定してください。")
		os.Exit(1)
	}

	// os.Argsを確認する。
	fmt.Println(os.Args)

	// sep := strings.Join(os.Args, "")
	sli := []string{os.Args}
	si, err := strconv.Atoi(sli)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println("err in converting")
	}

	fmt.Println(len(data) - i)

	// var splitcnt = len(data) / n

	// fmt.Println(splitcnt)

	// fmt.Println(data[len(data)-*i:])

}
