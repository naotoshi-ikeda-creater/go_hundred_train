package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//ファイルの読み取り
	f, err := os.Open("../data/neko.txt.mecab")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	//バッファリング
	r := bufio.NewReader(f)

	//mapの中にmapをつくるための宣言
	sents := make([][]map[string]string, 0)
	sent := make([]map[string]string, 0)

	//1行ずつ読んでいく
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		// fmt.Println(string(b))

		// EOS以外であれば、mapに割り当てていく
		if string(b) != "EOS" {
			//
			tmp := strings.Split(string(b), "\t")
			// fmt.Println(tmp)
			m := append(tmp[:1], strings.Split(tmp[1], ",")...)
			// fmt.Println(m)

			morpheme := make(map[string]string)

			morpheme["surface"] = m[0]
			morpheme["base"] = m[7]
			morpheme["pos"] = m[1]
			morpheme["pos1"] = m[2]

			sent = append(sent, morpheme)
		} else { // if we find "EOS", store sentence to sentences and initialize the sent
			if len(sent) > 0 { // EOSを見つけたら、そのぶんをsentsにappendして、sentを空の連想配列にする。
				//
				sents = append(sents, sent)
				sent = make([]map[string]string, 0)
			}
		}

	}

	fmt.Println(sents)

}
