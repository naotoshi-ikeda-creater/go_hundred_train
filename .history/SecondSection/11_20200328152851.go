package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ファイルの読み取りを開始します。")

	f, err := os.Open("hightemp.txt")
	if err != nil {
		fmt.Println("error")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() [
		fmt.Println(strings.Replace(scanner.Text,"\b"," ",-1))
	]
}
