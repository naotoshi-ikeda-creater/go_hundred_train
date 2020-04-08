package main

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {

	// jsonファイルを読み込み
	jsonString, err := readFile("../data/jawiki-country.json")
	if err != nil {
		panic(err)
	}
	var article []Article
	json.Unmarshal(jsonString, &article)

	for _, v := range article {
		fmt.Println(v.Text)
	}
}
