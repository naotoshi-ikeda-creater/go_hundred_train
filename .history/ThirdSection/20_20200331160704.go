package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Article struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {

	// jsonファイルを読み込み
	file, err := ioutil.ReadFile("../data/jawiki-country.json")
	if err != nil {
		panic(err)
	}
	var config Article
	json.Unmarshal(jsonString, &config)

	// jsonファイルの中身を表示
	fmt.Println(config.Text)
	fmt.Println(config.Title)
}
