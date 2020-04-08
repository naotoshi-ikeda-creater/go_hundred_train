package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type Article struct {
	Title string `json: "title"`
	Text  string `json: "text"`
}

func read(path string) []string {
	contents := []string{}

	//ファイルをOpenする

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer Close()

	//特定の文字まで読み込む
	//バッファリングする
	bu := bufio.NewReader(f)

	for {
		b, _ := bu.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		//Articleの型でaを宣言
		a := Article{}
		//その中にjsonのエンコードされたデータをaに格納する
		json.Unmarshal([]byte(b), &a)
		//そのaをcontentに入れていく
		contents = append(contents, a)

	}

	return contents

}

func main() {
	read("../data/jawiki-country.json")
}
