package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Article struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {
	articles := []Article{}

	f, err := os.Open("../data/jawiki-country.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		a := Article{}
		json.Unmarshal([]byte(b), &a)
		articles = append(articles, a)
	}

	var data string
	for _, article := range articles {
		if article.Title == "イギリス" {
			data = article.Text
		}
	}

	//正規表現初期化コンパイル
	reg := regexp.MustCompile(`ファイル:.*]`)
	//it returns a slice of all successive matches of the expression
	for _, v := range reg.FindAll([]byte(data), -1) {
		//byte型で返した値をstringへ変換し、セクション名の文字列だけを抜き出す
		fmt.Println(string(v[len("ファイル:") : len(v)-2]))
	}

}
