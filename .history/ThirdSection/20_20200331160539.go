package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Article struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {
	articles := []Article{}

	jsonString, err := ioutil.ReadFile("../data/jawiki-country.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// jsonファイルを読み込み
	file, err := ioutil.ReadFile("../data/jawiki-country.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)

	// jsonファイルの中身を表示
	fmt.Println(config.consumer_key)
	fmt.Println(config.consumer_secret)
}
