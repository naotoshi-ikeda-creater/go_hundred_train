package main

import (
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

	f, err := os.Open("../data/jawiki-country.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	for _, article := range articles {
		if article.Title == "イギリス" {
			fmt.Println(article.Text)
		}
	}

}
