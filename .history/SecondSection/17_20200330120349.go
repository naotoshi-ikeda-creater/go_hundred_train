package main

import (
	"fmt"
	"os"
)

func main() {
	// 引数の数を確認する。
	if len(os.Args) < 2 {
		fmt.Println("ERROR: 引数を指定してください。")
		os.Exit(1)
	}

	// os.Argsを確認する。
	fmt.Println(os.Args)

	// 個別引数へのアクセス。
	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("引数%d: %s\n", i, os.Args[i])
	}
}
