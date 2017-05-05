package smac

import (
	"os"
	"testing"
)

func init() {
	goPath := os.Getenv("GOPATH")

	wordFile := goPath + "/src/github.com/pierods/smac/allwords.txt"

	autoComplete, err := NewAutoCompleteF(wordFile, 0, 0)
	if err != nil {
		os.Exit(-1)
	}
	Ac = autoComplete
}

var Ac AutoComplete

func BenchmarkComplete(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Ac.Complete("chair")
	}
}
