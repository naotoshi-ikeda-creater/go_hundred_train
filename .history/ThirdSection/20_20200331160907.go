package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Article struct {
	text  string `json:"text"`
	title string `json:"title"`
}

func main() {

	// jsonファイルを読み込み
	jsonString, err := ioutil.ReadFile("../data/jawiki-country.json")
	if err != nil {
		panic(err)
	}
	var config Article
	json.Unmarshal(jsonString, &config)

	// jsonファイルの中身を表示
	fmt.Println(config.text)
	fmt.Println(config.title)
}
