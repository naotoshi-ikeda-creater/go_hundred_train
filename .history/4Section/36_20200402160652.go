package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}
func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}
func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

// sortedKeysは、map型で受け取り、mapのcountで並び替えられたstring型の配列で返す

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}

	sort.Sort(sm)

	return sm.s
}

func main() {
	f, err := os.Open("../data/neko.txt.mecab")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	sents := make([][]map[string]string, 0)
	sent := make([]map[string]string, 0)

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if string(b) != "EOS" {
			tmp := strings.Split(string(b), "\t")
			m := append(tmp[:1], strings.Split(tmp[1], ",")...)

			morpheme := make(map[string]string)

			morpheme["surface"] = m[0]
			morpheme["base"] = m[7]
			morpheme["pos"] = m[1]
			morpheme["pos1"] = m[2]

			sent = append(sent, morpheme)
		} else { // if we find "EOS", store sentence to sentences and initialize the sent
			if len(sent) > 0 { // for appearing "EOS" continuously
				sents = append(sents, sent)
				sent = make([]map[string]string, 0)
			}
		}

	}

	//
	freq := make(map[string]int)
	for _, sent := range sents {
		for _, m := range sent {
			freq[m["base"]]++
		}
	}

	ranking := sortedKeys(freq)
	for _, v := range ranking {
		fmt.Println(v, freq[v])
	}
}
