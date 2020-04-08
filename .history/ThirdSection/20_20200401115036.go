package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

/** JSONデコード用に構造体定義 */
type Person struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("../data/jawiki-country.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var persons []Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを表示
	for _, p := range persons {
		fmt.Printf("%d : %s\n", p.Text, p.Title)
	}
}
