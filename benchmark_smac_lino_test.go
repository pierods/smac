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

	ac, err := NewAutoCompleteLinoF(wordFile, 2, 0, 0)
	if err != nil {
		os.Exit(-1)
	}
	AcLino = ac

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

	prefixes := make(map[string]bool)
	for _, word := range Words {
		for i := 1; i < len(word); i++ {
			acc := word[:i]
			prefixes[acc] = true
		}
	}

	for k, _ := range prefixes {
		Prefixes = append(Prefixes, k)
	}

	Pl = len(Prefixes)
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
