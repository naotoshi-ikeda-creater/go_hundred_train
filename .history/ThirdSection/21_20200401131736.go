package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Article struct {
	Text     string `json:"text"`
	Title    string `json:"title"`
	Category string `json:"category`
}

func include(n string, data []Article) bool {
	for _, v := range data {
		if Article(n) == v {
			return true
		}
	}
	return false
}

func main() {
	contents := []Article{}

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
		contents = append(contents, a)
	}

	for _, article := range contents {
		if include("Category", article) {
			fmt.Println(article.Text)
		}
	}

}
