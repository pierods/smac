// Copyright Piero de Salvia.
// All Rights Reserved

package smac

import (
	"bufio"
	"math/rand"
	"os"
	"testing"
)

func init() {
	goPath := os.Getenv("GOPATH")

	wordFile := goPath + "/src/github.com/pierods/smac/demo/allwords.txt"

	benchAlphabet := "abcdefghijklmnopqrstuvwxyz1234567890'/&\""
	autoComplete, err := NewAutoCompleteF(benchAlphabet, wordFile, 0, 0)
	if err != nil {
		os.Exit(-1)
	}
	AcTrie = autoComplete

	f, err := os.Open(wordFile)

	defer f.Close()

	if err != nil {
		os.Exit(-1)
	}

	lineScanner := bufio.NewScanner(f)

	for lineScanner.Scan() {
		line := lineScanner.Text()
		Words = append(Words, line)
	}
	Wl = len(Words)

}

var AcTrie AutoCompleteTrie
var Words []string
var Prefixes []string
var Wl, Pl int

func BenchmarkCompleteTrie(b *testing.B) {

	for i := 0; i < b.N; i++ {
		w := Words[rand.Intn(Wl)]
		AcTrie.Complete(w)
	}
}
