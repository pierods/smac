// Copyright Piero de Salvia.
// All Rights Reserved

package smac

import (
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
}

var AcTrie AutoCompleteTrie

func BenchmarkCompleteTrie(b *testing.B) {

	for i := 0; i < b.N; i++ {
		AcTrie.Complete("chair")
	}
}
