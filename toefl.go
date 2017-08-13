package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func main() {
	word := random()
	fmt.Println(word)
}

type word struct {
	spelling string
	meanings []string
}

func (w word) String() string {
	var buf bytes.Buffer

	buf.WriteString(w.spelling)
	buf.WriteRune('\n')
	for _, meaning := range w.meanings {
		buf.WriteRune('\n')
		buf.WriteString(meaning)
	}

	return buf.String()
}

func random() word {
	words := load("words/wangyumei-toefl-words.txt")
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(words))
	return words[idx]
}

func load(path string) (words []word) {
	asset := MustAsset(path)
	for _, line := range strings.Split(string(asset), "\n") {
		line = strings.TrimFunc(line, unicode.IsSpace)
		if line == "" {
			continue
		}

		words = append(words, parse(line))
	}
	return
}

func parse(raw string) word {
	parts := strings.Split(raw, "#")
	spelling := parts[0]

	var meanings []string
	for _, meaning := range strings.Split(parts[1], ";") {
		meaning := strings.TrimFunc(meaning, unicode.IsSpace)
		meanings = append(meanings, meaning)
	}
	return word{spelling, meanings}
}
