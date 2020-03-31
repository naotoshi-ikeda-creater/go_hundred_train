package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("ファイルの読み込みを開始します。")

	f, err := os.Open("hightemp.txt")
	if err != nil {
		fmt.Printf("os.Open: %#v\n", err)
		return
	}

	//関数終了時に閉じる
	defer f.Close()

	//1行目用のファイルを作成する
	fi, err := os.Create("col1.txt")
	if err != nil {
		fmt.Printf("os.Open: %#v\n", err)
		return
	}

	//関数終了時に閉じる

	defer fi.Close()

	//2行目のファイルを作成する
	si, err := os.Create("col2.txt")
	if err != nil {
		fmt.Printf("os.Open: %#v\n", err)
		return
	}
	defer si.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		d := strings.Split(scanner.Text(), "\t")
		fi.WriteString(d[0] + "\n")		si.WriteString(d[1] + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v\n")
		return
	}
}
