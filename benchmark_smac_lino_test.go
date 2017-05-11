package smac

import (
	"os"
	"testing"
)

func init() {
	goPath := os.Getenv("GOPATH")

	wordFile := goPath + "/src/github.com/pierods/smac/demo/allwords.txt"

	ac, err := NewAutoCompleteLinoF(wordFile, 0, 0)
	if err != nil {
		os.Exit(-1)
	}
	AcLino = ac
}

var AcLino AutoCompleteLiNo

func BenchmarkCompleteLino(b *testing.B) {

	for i := 0; i < b.N; i++ {
		AcLino.Complete("chair")
	}
}
