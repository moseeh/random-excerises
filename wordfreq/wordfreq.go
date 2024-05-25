package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	filestring := string(file)
	filearr := strings.FieldsFunc(filestring, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	filemap := make(map[string]int)

	for _, v := range filearr {
		word := strings.ToLower(v)
		filemap[word]++
	}

	var wordFreqs []struct {
		word  string
		count int
	}
	for word, count := range filemap {
		wordFreqs = append(wordFreqs, struct {
			word  string
			count int
		}{word, count})
	}

	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].count > wordFreqs[j].count
	})

	for _, wf := range wordFreqs {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}
