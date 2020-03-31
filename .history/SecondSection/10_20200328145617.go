package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("ファイルの読み取りを開始します")

	line := 0

	//ファイルをOpenする
	f, err := os.Open("hightemp.txt")
	if err != nil {
		fmt.Println("error")
	}

	defer f.Close()

	//スキャナライブラリを作成
	scanner := bufio.NewScanner(f)

	//データを1行ずつ読み込む
	for scanner.Scan() {
		line++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v\n", err)
		return
	}

	//一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	//出力
	fmt.Println(string(b))
}
