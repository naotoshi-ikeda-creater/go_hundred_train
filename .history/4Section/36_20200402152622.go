package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Group struct {
	sur  string
	coun int
}

type SortTem []Group

//以下でインターフェースを満たす
//Len()
func (s SortTem) Len() int {
	return len(s)
}

//Less()
func (s SortTem) Less(i, j int) bool {
	return s[i].coun > s[j].coun
}

//Swap()
func (s SortTem) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	//ファイルの読み取り
	f, err := os.Open("../data/neko.txt.mecab")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	//バッファリング
	r := bufio.NewReader(f)

	//1行ずつ読んでいく
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		// fmt.Println(string(b))

		// EOS以外であれば、mapに割り当てていく
		if string(b) != "EOS" {
			//
			tmp := strings.Split(string(b), "\t")
			// fmt.Println(tmp)
			m := append(tmp[:1], strings.Split(tmp[1], ",")...)

			data := []rune{}
			for i, v := range m[0] {

				data = append(data, v)
			}

			fmt.Println(data)

			// 	data := []string{}
			// 	data = append(data)

			// 	counts := make(map[string]int)

			// 	for _, v := range m[0] {
			// 		counts[v]++
			// 	}

			// 	arr := []Group{}

			// 	for i, v := range counts {
			// 		tmp := Group{i, v}
			// 		arr = append(arr, tmp)
			// 	}

			// 	sort.Sort(SortTem(arr))

			// 	for _, v := range arr {
			// 		fmt.Printf("%v %v\n", v.pref, v.coun)
			// 	}

		}

	}
}
