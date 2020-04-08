package main

import (
	"bufio"
	"fmt"
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

	//関数終了時に閉じる
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
	fmt.Println("Line", line)

}
