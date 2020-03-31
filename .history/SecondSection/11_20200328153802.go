package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("ファイルの読み取りを開始します。")

	f, err := os.Open("hightemp.txt")
	if err != nil {
		fmt.Println("error")
	}

	//関数終了時に閉じる
	defer f.Close()

	//スキャナライブラリを作成
	scanner := bufio.NewScanner(f)

	//1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		fmt.Println(strings.Replace(scanner.Text(), "\b", " ", -1))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v\n", err)
		return
	}
}
