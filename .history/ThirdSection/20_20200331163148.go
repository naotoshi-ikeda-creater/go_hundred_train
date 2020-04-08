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

func readFile(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	return bytes, err
}

func main() {

	// jsonファイルを読み込み
	jsonString, err := readFile("../data/jawiki-country.json")
	if err != nil {
		panic(err)
	}
	var article Article
	json.Unmarshal(jsonString, &config)

	// jsonファイルの中身を表示
	fmt.Println(config.Text)
	fmt.Println(config.Title)
}
