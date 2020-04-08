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

	c := new(Config)

	err = json.Unmarshal(jsonString, c)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	for _, article := range articles {
		if article.Title == "イギリス" {
			fmt.Println(article.Text)
		}
	}

}
