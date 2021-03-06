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

	ac, err := NewAutoCompleteLinoF(wordFile, 4, 10, 90)
	if err != nil {
		os.Exit(-1)
	}
	AcLino = ac
}

var AcLino AutoCompleteLiNo

func BenchmarkLinoCompleteWords(b *testing.B) {

	var r []string

	for i := 0; i < b.N; i++ {
		w := testWords[rand.Intn(wordsInTestData)]
		r, _ = AcLino.Complete(w)
	}

	result = r
}

func BenchmarkLinoPrefixes(b *testing.B) {

	var r []string

	for i := 0; i < b.N; i++ {
		p := testPrefixes[rand.Intn(prefixesInTestData)]
		r, _ = AcLino.Complete(p)
	}

	result = r
}

func BenchmarkLinoMix(b *testing.B) {

	var r []string

	flip := false

	for i := 0; i < b.N; i++ {
		if flip {
			p := testPrefixes[rand.Intn(prefixesInTestData)]
			r, _ = AcLino.Complete(p)

		} else {
			w := testWords[rand.Intn(wordsInTestData)]
			r, _ = AcLino.Complete(w)
		}
		flip = !flip
	}

	result = r
}

var result []string
