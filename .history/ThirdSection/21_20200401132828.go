package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func member(n string, data []string) bool {
	for _, v := range data {
		if n == v {
			return true
		}
	}
	return false
}

type Article struct {
	Text  string `json:"text"`
	Title string `json:"title"`
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

	data := []string{}
	for _, article := range contents {

		data = append(data, article.Text)

	}

	answer := []string{}

	for _, v := range data {

		if  {
			answer = append(answer, v)

		}

	}

	fmt.Println(answer)
}
