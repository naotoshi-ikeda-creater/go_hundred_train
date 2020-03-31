package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Cluster struct {
	pref string
	coun int
}

type SortTemp []Cluster

//以下でインターフェースを満たす
//Len()
func (s SortTemp) Len() int {
	return len(s)
}

//Less()
func (s SortTemp) Less(i, j int) bool {
	return s[i].coun > s[j].coun
}

//Swap()
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
	//データをdataへ読み込む
	data := read("col1.txt")

	counts := make(map[string]int)

	for _, v := range data {
		counts[v]++
	}

	arr := []Cluster{}

	for i, v := range counts {
		tmp := Cluster{i, v}
		arr = append(arr, tmp)
	}

	sort.Sort(SortTemp(arr))

	for _, v := range arr {
		fmt.Printf("%v %v\n", v.pref, v.coun)
	}

}
