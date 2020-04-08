package main

import (
	"fmt"
	"os"
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

	for _, article := range articles {
		if article.Title == "イギリス" {
			fmt.Println(article.Text)
		}
	}

}
