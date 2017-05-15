// Copyright Piero de Salvia.
// All Rights Reserved

package smac

import (
	"math/rand"
	"os"
	"testing"
)

func init() {
	initBenchmark()
	goPath := os.Getenv("GOPATH")

	wordFile := goPath + "/src/github.com/pierods/smac/demo/allwords.txt"

	benchAlphabet := "abcdefghijklmnopqrstuvwxyz1234567890'/&\""
	autoComplete, err := NewAutoCompleteF(benchAlphabet, wordFile, 0, 0)
	if err != nil {
		os.Exit(-1)
	}
	AcTrie = autoComplete

}

var AcTrie AutoCompleteTrie

func BenchmarkTrieCompleteWords(b *testing.B) {

	for i := 0; i < b.N; i++ {
		w := Words[rand.Intn(Wl)]
		AcTrie.Complete(w)
	}
}

func BenchmarkTriePrefixes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p := Prefixes[rand.Intn(Pl)]
		AcTrie.Complete(p)
	}
}
