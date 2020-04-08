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

	res := make([]string, 0)
	for _, sent := range sents {
		//iの初期値を0にし、for i < len(sent)-1 で index out of []のエラーを防ぐ
		i := 0
		for i < len(sent)-1 {
			if sent[i]["pos"] == "名詞" {
				j := 1
				tmp := sent[i]["surface"]
				for sent[i+j]["pos"] == "名詞" {
					tmp += sent[i+j]["surface"]
					j++
					//上記と同じ
					if i+j > len(sent)-1 {
						break
					}
				}
				//iの値を正常な順に戻す
				i += j
				res = append(res, tmp)
			} else {
				//次へ進む
				i++
			}
		}
	}

	fmt.Println(res)
}
