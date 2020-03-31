package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ファイルの読み込みを開始します。")

	f, err := os.Open("hightemp.txt")
	if err != nil {
		fmt.Println("error")
	}

	//関数終了時に閉じる
	defer f.Close()
}
