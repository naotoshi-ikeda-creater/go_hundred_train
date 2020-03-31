package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fi, si := []string{}, []string{}

	read := func(filename string) []string {
		//読み込むファイルを開く
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

	for _, v ;= range read("hightemp.txt") {
		rp := strings.Fields(v)
		fi = append(fi, rp[0])
		si = append(si, rp[1])
	}

	write := func (filename string, data []string) {
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		w := bufio.NewWriter(f)
		for _, v := range  data {
			fmt.Fprintf(w,v,"\n")
		}

	}
}
