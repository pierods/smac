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

	for i := 0; i < b.N; i++ {
		w := Words[rand.Intn(Wl)]
		AcLino.Complete(w)
	}
}

func BenchmarkLinoPrefixes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p := Prefixes[rand.Intn(Pl)]
		AcLino.Complete(p)
	}
}

func BenchmarkLinoMix(b *testing.B) {

	flip := false

	for i := 0; i < b.N; i++ {
		if flip {
			p := Prefixes[rand.Intn(Pl)]
			AcLino.Complete(p)

		} else {
			w := Words[rand.Intn(Wl)]
			AcLino.Complete(w)
		}
		flip = !flip
	}
}
