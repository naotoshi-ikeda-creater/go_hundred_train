package main

import (
	"bufio"
	"fmt"
	"os"
)

type cluster struct {
	pref string
	city string
	temp float64
	date string
}

type SortTemp []cluster

func (s SortTemp) Len() int {
	return len(s)
}

func (s SortTemp) Less(i, j int) bool {
	return s[i].temp < s[j].temp
}

func (s SortTemp) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

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

}
