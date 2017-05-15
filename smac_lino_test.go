package smac

import (
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

	t.Log("Given the need to test the lino un learn coherence")
	{
		autoComplete, _ := NewAutoCompleteLinoS(words, 2, 0, 0)
		autoComplete.UnLearn("aaaa")
		newWords := []string{"aabc", "abc", "bbb", "v", "vvv", "vvvaaa"}
		for i, word := range newWords[:len(newWords)-1] {
			nextWord := autoComplete.wordMap[word].next
			if nextWord != newWords[i+1] {
				t.Fatal("Should be able to rebuild a linked list of dictionary words", ballotX)
			}
		}
		t.Log("Should be able to rebuild a linked list of dictionary words", checkMark)

		if autoComplete.head != "aabc" {
			t.Fatal("Should be able to correctly reassign the head of an autocompleter", ballotX)
		}
		t.Log("Should be able to correctly reassign the head of an autocompleter", checkMark)
		prefix2["aa"] = "aabc"
		prefix2["a"] = "aabc"
		if !reflect.DeepEqual(autoComplete.prefixMap, prefix2) {
			t.Log(autoComplete.prefixMap)
			t.Fatal("Should be able to rebuild the prefix map of an autocompleter on a head removal", ballotX)
		}
		t.Log("Should be able to rebuild the prefix map of an autocompleter on a head removal", checkMark)

		autoComplete.UnLearn("vvvaaa")
		newWords = []string{"aabc", "abc", "bbb", "v", "vvv"}
		for i, word := range newWords[:len(newWords)-1] {
			nextWord := autoComplete.wordMap[word].next
			if nextWord != newWords[i+1] {
				t.Fatal("Should be able to rebuild a linked list of dictionary words", ballotX)
			}
		}
		t.Log("Should be able to rebuild a linked list of dictionary words", checkMark)

		if autoComplete.tail != "vvv" {
			t.Fatal("Should be able to correctly reassign the tail of an autocompleter", ballotX)
		}
		t.Log("Should be able to correctly reassign the tail of an autocompleter", checkMark)

		prefix2["vv"] = "vvv"
		if !reflect.DeepEqual(autoComplete.prefixMap, prefix2) {
			t.Fatal("Should be able to rebuild the prefix map of an autocompleter on a tail removal", ballotX)
		}
		t.Log("Should be able to rebuild the prefix map of an autocompleter on a tail removal", checkMark)

		autoComplete.UnLearn("bbb")
		newWords = []string{"aabc", "abc", "v", "vvv"}
		for i, word := range newWords[:len(newWords)-1] {
			nextWord := autoComplete.wordMap[word].next
			if nextWord != newWords[i+1] {
				t.Fatal("Should be able to rebuild a linked list of dictionary words", ballotX)
			}
		}
		t.Log("Should be able to rebuild a linked list of dictionary words", checkMark)
		delete(prefix2, "b")
		delete(prefix2, "bb")
		if !reflect.DeepEqual(autoComplete.prefixMap, prefix2) {
			t.Log(autoComplete.prefixMap)
			t.Log(prefix2)
			t.Fatal("Should be able to rebuild the prefix map of an autocompleter on prefix word removal", ballotX)
		}
		t.Log("Should be able to rebuild the prefix map of an autocompleter on prefix word removal", checkMark)

		initTestVals()
		autoComplete, _ = NewAutoCompleteLinoS(words, 2, 0, 0)
		autoComplete.UnLearn("aabc")

		newWords = []string{"aaaa", "abc", "bbb", "v", "vvv", "vvvaaa"}
		for i, word := range newWords[:len(newWords)-1] {
			nextWord := autoComplete.wordMap[word].next
			if nextWord != newWords[i+1] {
				t.Fatal("Should be able to rebuild a linked list of dictionary words", ballotX)
			}
		}
		t.Log("Should be able to rebuild a linked list of dictionary words", checkMark)
		if !reflect.DeepEqual(autoComplete.prefixMap, prefix2) {
			t.Fatal("Should be able to maintain prefix map of an autocompleter on a non-prefix word removal", ballotX)
		}
		t.Log("Should be able to maintain the prefix map of an autocompleter on a non-prefix word removal", checkMark)
	}
}

func Test_LinoSaveAndRetrieve(t *testing.T) {
	tempDir := os.TempDir()
	tempFile, err := ioutil.TempFile(tempDir, "smac")
	if err != nil {
		t.Fatal(err)
	}
	fName := tempFile.Name()

	initTestVals()

	autoComplete, _ := NewAutoCompleteLinoS(words, 2, 0, 0)
	autoComplete.Accept("aaaa")
	autoComplete.Learn("ddd")
	autoComplete.Learn("eee")
	autoComplete.Accept("eee")
	autoComplete.UnLearn("vvv")

	t.Log("Given the need to test the save/retrieve feature")
	{
		err = autoComplete.Save(tempFile.Name())

		if err != nil {
			t.Fatal("Should be able to save words to a file", ballotX)
		}
		t.Log("Should be able to save words to a file", checkMark)

		f, err := os.Open(fName)
		defer f.Close()
		if err != nil {
			t.Fatal(err)
		}
		dec := gob.NewDecoder(f)
		readWords := make(map[string]int)
		for {
			var wA wordAccepts
			if err = dec.Decode(&wA); err == io.EOF {
				break
			} else if err != nil {
				t.Fatal("Should be able to read a smac file")
			}
			readWords[wA.Word] = wA.Accepts
		}

		accepts, exists := readWords["aaaa"]

		if !exists {
			t.Fatal("Should be able to persist an accepted word", ballotX)
		}
		t.Log("Should be able to perstist an accepted word", checkMark)

		if accepts != 1 {
			t.Fatal("Should be able to persist an accepted word", ballotX)
		}
		t.Log("Should be able to perstist an accepted word", checkMark)

		accepts, exists = readWords["ddd"]

		if !exists {
			t.Fatal("Should be able to persist a learnt word", ballotX)
		}
		t.Log("Should be able to persist a learnt word", checkMark)

		accepts, exists = readWords["eee"]

		if !exists {
			t.Fatal("Should be able to persist a learnt word", ballotX)
		}
		t.Log("Should be able to persist a learnt word", checkMark)

		if accepts != 1 {
			t.Fatal("Should be able to persist a learnt and accepted word", ballotX)
		}
		t.Log("Should be able to persist a learnt and accepted word", checkMark)

		accepts, exists = readWords["vvv"]

		if !exists {
			t.Fatal("Should be able to persist an unlearnt word", ballotX)
		}
		t.Log("Should be able to persist an unlearnt word", checkMark)

		if accepts != -1 {
			t.Fatal("Should be able to persist a learnt and accepted word", ballotX)
		}
		t.Log("Should be able to persist a learnt and accepted word", checkMark)

		autoComplete, _ = NewAutoCompleteLinoS(words, 2, 0, 0)
		err = autoComplete.Retrieve(fName)
		if err != nil {
			t.Fatal(err)
		}
		/*
			ac, _ := autoComplete.Complete("aaaa")
			if !reflect.DeepEqual(ac, []string{"aaabbb", "aaa"}) {
				t.Fatal("Should be able to get back from retrieve an accepted word", ballotX)
			}
			t.Log("Should be able to get back from retrieve an accepted word", checkMark)
			ac, _ = autoComplete.Complete("ddd")
			if !reflect.DeepEqual(ac, []string{"ddd"}) {
				t.Fatal("Should be able to get back from retrieve a learned word", ballotX)
			}
			t.Log("Should be able to get back from retrieve a learned word", checkMark)
			ac, _ = autoComplete.Complete("ccc")
			if !reflect.DeepEqual(ac, []string{}) {
				t.Fatal("Should be able to erase from retrieve a deleted word", ballotX)
			}
			t.Log("Should be able to erase from retrieve a deleted word", checkMark)
		*/
	}
}

func Example_NewAutoCompleteLinoS() {

	words := []string{"chair", "chairman", "chairperson", "chairwoman"}
	autoComplete, err := NewAutoCompleteLinoS(words, 2, 0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	completes, err := autoComplete.Complete("chair")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(completes)
}
