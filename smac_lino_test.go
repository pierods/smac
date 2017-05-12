package smac

import (
	"reflect"
	"sort"
	"testing"
)

var (
	words   []string
	prefix0 map[string]string
	prefix1 map[string]string
	prefix2 map[string]string
	prefix3 map[string]string
)

func initTestVals() {

	words = []string{"aaaa", "aabc", "abc", "bbb", "v", "vvv", "vvvaaa"}
	prefix0 = make(map[string]string)
	prefix1 = make(map[string]string)
	prefix2 = make(map[string]string)
	prefix3 = make(map[string]string)

	prefix1["a"] = "aaaa"
	prefix1["b"] = "bbb"
	prefix1["v"] = "v"

	for k, v := range prefix1 {
		prefix2[k] = v
	}

	prefix2["aa"] = "aaaa"
	prefix2["ab"] = "abc"
	prefix2["bb"] = "bbb"
	prefix2["vv"] = "vvv"

	for k, v := range prefix2 {
		prefix3[k] = v
	}

	prefix3["aaa"] = "aaaa"
	prefix3["aab"] = "aabc"
	prefix3["abc"] = "abc"
	prefix3["bbb"] = "bbb"
	prefix3["vvv"] = "vvv"

	sort.Strings(words)
}

func Test_LinoConstruction(t *testing.T) {

	initTestVals()

	t.Log("Given the need to test the prefix map creation")
	{

		prefixes := makePrefixMap(words, 0)
		if !reflect.DeepEqual(prefixes, prefix0) {
			t.Log("Should be able to create a 0 len prefix map", ballotX)
		}
		t.Log("Should be able to create a 0 len prefix map", checkMark)

		prefixes = makePrefixMap(words, 1)

		if !reflect.DeepEqual(prefixes, prefix1) {
			t.Fatal("Should be able to create a 1 len prefix map", ballotX)
		}
		t.Log("Should be able to create a 1 len prefix map", checkMark)

		prefixes = makePrefixMap(words, 2)

		if !reflect.DeepEqual(prefixes, prefix2) {
			t.Fatal("Should be able to create a 2 len prefix map", ballotX)
		}
		t.Log("Should be able to create a 2 len prefix map", checkMark)

		prefixes = makePrefixMap(words, 3)

		if !reflect.DeepEqual(prefixes, prefix3) {
			t.Fatal("Should be able to create a 3 len prefix map", ballotX)
		}
		t.Log("Should be able to create a 3 len prefix map", checkMark)
	}

	t.Log("Given the need to test the lino creation")
	{
		words := []string{"aaaa", "aabc", "abc", "bbb", "v", "vvv", "vvvaaa"}
		autoComplete, _ := NewAutoCompleteLinoS(words, 0, 0, 0)

		for i, word := range words[:len(words)-1] {
			nextWord := autoComplete.wordMap[word].next
			if nextWord != words[i+1] {
				t.Fatal("Should be able to build a linked list of dictionary words", ballotX)
			}
		}
		t.Log("Should be able to build a linked list of dictionary words", checkMark)
		if autoComplete.head != "aaaa" {
			t.Fatal("Should be able to correctly initialize the head of an autocompleter", ballotX)
		}
		t.Log("Should be able to correctly initialize the head of an autocompleter", checkMark)
		if autoComplete.tail != "vvvaaa" {
			t.Fatal("Should be able to correctly initialize the tail of an autocompleter", ballotX)
		}
		t.Log("Should be able to correctly initialize the tail of an autocompleter", checkMark)
	}
}

func Test_LinoFindPreviousWord(t *testing.T) {

	initTestVals()
	t.Log("Given the need to test the lino navigation function")
	{
		autoComplete, _ := NewAutoCompleteLinoS(words, 2, 0, 0)
		if autoComplete.findPreviousWord("vvv") != "v" {
			t.Fatal("Should be able to correctly navigate a lino on a prefix word", ballotX)
		}
		t.Log("Should be able to correctly navigate a lino on a prefix word", checkMark)

		if autoComplete.findPreviousWord("bba") != "abc" {
			t.Fatal("Should be able to correctly navigate a lino on a pre-prefix word", ballotX)
		}
		t.Log("Should be able to correctly navigate a lino on a pre-prefix word", checkMark)
		if autoComplete.findPreviousWord("bbc") != "bbb" {
			t.Fatal("Should be able to correctly navigate a lino on a post-prefix word", ballotX)
		}
		t.Log("Should be able to correctly navigate a lino on a post-prefix word", checkMark)
		if autoComplete.findPreviousWord("t") != "bbb" {
			t.Fatal("Should be able to correctly navigate a lino on a non-prefix word", ballotX)
		}
		t.Log("Should be able to correctly navigate a lino on a non-prefix word", checkMark)
		if autoComplete.findPreviousWord("a") != autoComplete.head {
			t.Fatal("Should be able to correctly navigate a lino on a pre-head word", ballotX)
		}
		t.Log("Should be able to correctly navigate a lino on a pre-head word", checkMark)
		if autoComplete.findPreviousWord("z") != autoComplete.tail {
			t.Fatal("Should be able to correctly navigate a lino on a post-tail word", ballotX)
		}
		t.Log("Should be able to correctly navigate a lino on a post-tail word", checkMark)

	}
}

func Test_LinoLearnCoherence(t *testing.T) {

	initTestVals()

	t.Log("Given the need to test the lino learn coherence")
	{
		autoComplete, _ := NewAutoCompleteLinoS(words, 2, 0, 0)
		autoComplete.Learn("bba")
		newWords := []string(words)
		newWords = append(newWords, "bba")
		sort.Strings(newWords)

		for i, word := range newWords[:len(newWords)-1] {
			nextWord := autoComplete.wordMap[word].next
			if nextWord != newWords[i+1] {
				t.Fatal("Should be able to rebuild a linked list of dictionary words", ballotX)
			}
		}
		t.Log("Should be able to rebuild a linked list of dictionary words", checkMark)

		prefix2["bb"] = "bba"
		prefix2["b"] = "bba"
		if !reflect.DeepEqual(autoComplete.prefixMap, prefix2) {
			t.Fatal("Should be able to rebuild the prefix map of an autocompleter", ballotX)
		}
		t.Log("Should be able to rebuild the prefix map of an autocompleter", checkMark)

		autoComplete.Learn("zzz")
		newWords = append(newWords, "zzz")
		sort.Strings(newWords)
		for i, word := range newWords[:len(newWords)-1] {
			nextWord := autoComplete.wordMap[word].next
			if nextWord != newWords[i+1] {
				t.Fatal("Should be able to rebuild a linked list of dictionary words on a tail word", ballotX)
			}
		}
		t.Log("Should be able to rebuild a linked list of dictionary words on a tail word", checkMark)

		if autoComplete.tail != "zzz" {
			t.Fatal("Should be able to correctly re-initialize the tail of an autocompleter", ballotX)
		}
		t.Log("Should be able to correctly re-initialize the tail of an autocompleter", checkMark)

		prefix2["zz"] = "zzz"
		prefix2["z"] = "zzz"
		if !reflect.DeepEqual(autoComplete.prefixMap, prefix2) {
			t.Fatal("Should be able to rebuild the prefix map of an autocompleter on a tail insert", ballotX)
		}
		t.Log("Should be able to rebuild the prefix map of an autocompleter on a tail insert", checkMark)

		autoComplete.Learn("a")
		newWords = append(newWords, "a")
		sort.Strings(newWords)
		for i, word := range newWords[:len(newWords)-1] {
			nextWord := autoComplete.wordMap[word].next
			if nextWord != newWords[i+1] {
				t.Fatal("Should be able to rebuild a linked list of dictionary words on a head word", ballotX)
			}
		}
		t.Log("Should be able to rebuild a linked list of dictionary words on a tail word", checkMark)
		if autoComplete.head != "a" {
			t.Fatal("Should be able to correctly initialize the head of an autocompleter", ballotX)
		}
		t.Log("Should be able to correctly re-initialize the head of an autocompleter", checkMark)

		prefix2["a"] = "a"
		if !reflect.DeepEqual(autoComplete.prefixMap, prefix2) {
			t.Fatal("Should be able to rebuild the prefix map of an autocompleter on a head insert", ballotX)
		}
		t.Log("Should be able to rebuild the prefix map of an autocompleter on a head insert", checkMark)

	}
}

func Test_LinoUnLearnCoherence(t *testing.T) {

	initTestVals()

	t.Log("Given the need to test the lino learn coherence")
	{
	}
}
