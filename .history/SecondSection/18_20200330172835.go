package main

import (
	"bufio"
	"fmt"
	"os"
)

//型リテラルを基にしてSortTempという型を定義
type Cluster struct {
	pref string
	city string
	temp float64
	date string
}

構造体を

type SortTemp []Cluster

//Len()の型
func (s SortTemp) Len() int {
	return len(s)
}

//Less()の型
func (s SortTemp) Less(i, j int) bool {
	return s[i].temp < s[j].temp
}

//Swap()の型
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
	data := read("../data/hightemp.txt")

	arr := []Cluster{}

	for _, v:= range data {
		rep := strings.Fields(v)
		stoi,_ := strconv.ParseFloat(rep[2],64)
		newdata := Cluster{rep[0],rep[1], stoi,rep[3]}
		arr = append(arr, newdata)

	}

	sort.Sort(SortTemp(arr))

	for _, v:= range arr {
		fmt.Printf("%v\t%v\t%v\t%v\n", v.pref, v.city,v.temp,v.date)
	}
}
