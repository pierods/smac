// Copyright Piero de Salvia.
// All Rights Reserved

package smac

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func initBenchmark() {
	goPath := os.Getenv("GOPATH")

	wordFile := goPath + "/src/github.com/pierods/smac/demo/allwords.txt"

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

	for k := range prefixes {
		Prefixes = append(Prefixes, k)
	}

	Pl = len(Prefixes)

}

var Words []string
var Prefixes []string
var Wl, Pl int

func TestTrieSOLILI(t *testing.T) {

	t.Log("Given the need to test a sorted linked list")
	{
		t.Log("When initializing a solili")
		{
			list := sOLILI{}
			slice := list.flush()
			if !reflect.DeepEqual(slice, []string{}) {
				t.Fatal("Should be able to correctly initialize an empty solili", ballotX)
			}
			t.Log("Should be able to correctly initialize an empty solili", checkMark)
		}
		t.Log("When operating on a solili")
		{
			list := sOLILI{}
			list.insert("aaa", 0)
			slice := list.flush()
			if !reflect.DeepEqual(slice, []string{"aaa"}) {
				t.Fatal("Should be able to correctly add an element to a solili", ballotX)
			}
			t.Log("Should be able to correctly add an element to a solili", checkMark)
			list.insert("bbb", 0)
			slice = list.flush()
			if !reflect.DeepEqual(slice, []string{"aaa", "bbb"}) {
				t.Fatal("Should be able to correctly add an element to the back of a solili", ballotX)
			}
			t.Log("Should be able to correctly add an element to the back of a solili", checkMark)
			list.insert("jjj", 100)
			slice = list.flush()
			if !reflect.DeepEqual(slice, []string{"jjj", "aaa", "bbb"}) {
				t.Fatal("Should be able to correctly add an element to the front of a solili", ballotX)
			}
			t.Log("Should be able to correctly add an element to the front of a solili", checkMark)
			list.insert("kkk", 90)
			slice = list.flush()
			if !reflect.DeepEqual(slice, []string{"jjj", "kkk", "aaa", "bbb"}) {
				t.Fatal("Should be able to correctly add an element in the middle of a solili", ballotX)
			}
			t.Log("Should be able to correctly add an element in the middle of a solili", checkMark)
			list.insert("lll", 100)
			slice = list.flush()
			if !reflect.DeepEqual(slice, []string{"jjj", "lll", "kkk", "aaa", "bbb"}) {
				t.Fatal("Should be able to maintain arrival order for same-weight elements (at head)", ballotX)
			}
			t.Log("Should be able to maintain arrival order for same-weight elements (at head)", checkMark)
			list.insert("mmm", 90)
			slice = list.flush()
			if !reflect.DeepEqual(slice, []string{"jjj", "lll", "kkk", "mmm", "aaa", "bbb"}) {
				t.Fatal("Should be able to maintain arrival order for same-weight elements", ballotX)
			}
			t.Log("Should be able to maintain arrival order for same-weight elements", checkMark)

			list = sOLILI{}
			list.insert("1", 0)
			list.insert("2", 0)
			list.insert("3", 0)
			list.insert("4", 0)
			list.insert("5", 0)
			list.insert("6", 0)
			list.insert("7", 0)
			list.insert("8", 0)
			list.insert("8", 0)

			slice = list.flushL(5)
			if len(slice) != 5 {
				t.Fatal("Should be able to limit flushing on a lili", ballotX)
			}
			t.Log("Should be able to limit flushing on a lili", checkMark)
		}
	}
}
