// Copyright Piero de Salvia.
// All Rights Reserved

package smac

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"testing"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func Test_SOLILI(t *testing.T) {

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
			list.insert("aaa", 0)
			slice = list.flush()
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
		}
	}
}

func Test_FIFO(t *testing.T) {

	var fifo fIFO

	t.Log("Given the need to test a FIFO")
	{
		t.Log("When initializing a FIFO")
		{
			if fifo.size() != 0 {
				t.Fatal("Should be able to initialize an empty FIFO", ballotX)
			}
			t.Log("Shoulf be able to initialize an empty FIFO", checkMark)
		}

		t.Log("When growing/shrinking a FIFO")
		{
			rarray := []rune("a")
			a := branch{
				node:   nil,
				parent: &rarray,
			}

			fifo.add(a)

			if fifo.size() != 1 {
				t.Fatal("Should be able to grow a FIFO by 1", ballotX)
			}
			t.Log("Should be able to grow a FIFO by 1", checkMark)
			newA := fifo.remove()
			if !reflect.DeepEqual(newA, a) {
				t.Fatal("Should be able to retrieve an element from a FIFO", ballotX)
			}
			t.Log("Should be able to retrieve an element from a FIFO", checkMark)
			rarray = []rune("b")
			b := branch{
				node:   nil,
				parent: &rarray,
			}
			rarray = []rune("c")
			c := branch{
				node:   nil,
				parent: &rarray,
			}
			fifo.add(a)
			fifo.add(b)
			fifo.add(c)
			if fifo.size() != 3 {
				t.Fatal("Should be able to grow a fifo by 3", ballotX)
			}
			t.Log("Should be able to grow a fifo by 3", checkMark)
			el1 := fifo.remove()
			el2 := fifo.remove()
			el3 := fifo.remove()
			if !reflect.DeepEqual(el1, a) || !reflect.DeepEqual(el2, b) || !reflect.DeepEqual(el3, c) {
				t.Fatal("Should be able to shrink a fifo in order", ballotX)
			}
			t.Log("Should be able to shrink a fifo in order", checkMark)
			if fifo.size() != 0 {
				t.Fatal("Should be able to shrink a fifo to zero", ballotX)
			}
			t.Log("Should be able to shrink a fifo to zero", checkMark)
		}
	}
}

func Test_LIFO(t *testing.T) {

	var lifo lIFO

	t.Log("Given the need to test a LIFO")
	{
		t.Log("When initializing a LIFO")
		{
			if lifo.size() != 0 {
				t.Fatal("Should be able to initialize an empty LIFO", ballotX)
			}
			t.Log("Shoulf be able to initialize an empty LIFO", checkMark)
		}

		t.Log("When growing/shrinking a LIFO")
		{
			tNp := &trieNode{}

			lifo.push(tNp)

			if lifo.size() != 1 {
				t.Fatal("Should be able to grow a LIFO by 1", ballotX)
			}
			t.Log("Should be able to grow a LIFO by 1", checkMark)
			newTNp := lifo.pop()
			if newTNp != tNp {
				t.Fatal("Should be able to pop an element from a LIFO", ballotX)
			}
			t.Log("Should be able to pop an element from a LIFO", checkMark)
			tNp2 := &trieNode{}
			tNp3 := &trieNode{}
			lifo.push(tNp)
			lifo.push(tNp2)
			lifo.push(tNp3)
			if lifo.size() != 3 {
				t.Fatal("Should be able to grow a lifo by 3", ballotX)
			}
			t.Log("Should be able to grow a lifo by 3", checkMark)
			el3 := lifo.pop()
			el2 := lifo.pop()
			el1 := lifo.pop()
			if el1 != tNp || el2 != tNp2 || el3 != tNp3 {
				t.Fatal("Should be able to shrink a lifo in order", ballotX)
			}
			t.Log("Should be able to shrink a lifo in order", checkMark)
			if lifo.size() != 0 {
				t.Fatal("Should be able to shrink a lifo to zero", ballotX)
			}
			t.Log("Should be able to shrink a lifo to zero", checkMark)
		}
	}
}

func Test_RunesToInts(t *testing.T) {

	words := []string{"aaa", "aaaa", "aaab", "aaac"}
	autoComplete, _ := NewAutoCompleteS(alphabet, words, 0, 0)

	t.Log("Given the need to test the runesToInts() function")
	{
		ints, err := autoComplete.runesToInts("")
		if err != nil {
			t.Fatal("Should be able to convert an empty string", ballotX)
		}
		if len(ints) != 0 {
			t.Fatal("Should be able to convert an empty string", ballotX)
		}
		t.Log("Should be able to convert an empty string", checkMark)
		ints, err = autoComplete.runesToInts("a")
		if err != nil {
			t.Fatal(err)
		}
		if len(ints) != 1 || ints[0] != 97 {
			t.Fatal("Should be able to convert a 1-len string", ballotX)
		}
		t.Log("Should be able to convert a 1-len string", checkMark)
		ints, err = autoComplete.runesToInts("abc")
		if err != nil {
			t.Fatal(err)
		}
		if len(ints) != 3 || !reflect.DeepEqual(ints, []int{97, 98, 99}) {
			t.Fatal("Should be able to convert a 3-len string", ballotX)
		}
		t.Log("Should be able to convert a 3-len string", checkMark)
		ints, err = autoComplete.runesToInts("漢")
		if err == nil {
			t.Fatal("Should be able to recognize a non-init-alphabet string", ballotX)
		}
		t.Log("Should be able to recognize a non-init-alphabet string", checkMark)
	}

}

func Test_TrieConstruction(t *testing.T) {
	words := []string{"abc"}

	autoComplete, _ := NewAutoCompleteS(alphabet, words, 0, 0)

	t.Log("Given the need to test the putIter() function")
	{
		a := autoComplete.root.links[0]
		if a == nil {
			t.Fatal("Should be able to insert first character of a word in tree", ballotX)
		}
		if a.intRune != 97 {
			t.Fatal("Should be able to insert first character of a word in tree", ballotX)
		}
		if a.isWord {
			t.Fatal("Should be able to insert first character of a word in tree", ballotX)
		}
		t.Log("Should be able to insert first character of a word in tree", checkMark)
		b := a.links[1]
		if b.intRune != 98 {
			t.Fatal("Should be able to insert second character of a word in tree", ballotX)
		}
		if b.isWord {
			t.Fatal("Should be able to insert second character of a word in tree", ballotX)
		}
		t.Log("Should be able to insert second character of a word in tree", checkMark)
		c := b.links[2]
		if !c.isWord {
			t.Fatal("Should be able to insert a word in tree", ballotX)
		}
		t.Log("Should be able to insert a word in tree", checkMark)
	}
	t.Log("Given the need to test a non-ASCII alphabet")
	{
		rAlphabet := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
		words = []string{"абзац"}
		autoComplete, err := NewAutoCompleteS(rAlphabet, words, 0, 0)
		if err != nil {
			t.Fatal("Should be able to instantiate an autcomplete on a non-ASCII alphabet", ballotX)
		}
		t.Log("Should be able to instantiate an autcomplete on a non-ASCII alphabet", checkMark)
		russianA := autoComplete.root.links[0]
		if russianA == nil {
			t.Fatal("Should be able to insert first character of a non-ASCII word in tree", ballotX)
		}
		if russianA.intRune != 1072 {
			t.Fatal("Should be able to insert first character of a non-ASCII word in tree", ballotX)
		}
		completes, err := autoComplete.Complete("абзац")
		if err != nil {
			t.Fatal("Should be able to complete on a non-ASCII word", ballotX)
		}
		if !reflect.DeepEqual(completes, []string{"абзац"}) {
			t.Fatal("Should be able to complete on a non-ASCII word", ballotX)
		}
		t.Log("Should be able to complete on a non-ASCII word", checkMark)
	}
}

func Test_Completion(t *testing.T) {

	words := []string{"aaa", "aaab", "aaac", "aaad", "abbbbb"}
	autoComplete, _ := NewAutoCompleteS(alphabet, words, 0, 0)
	ac1, _ := autoComplete.Complete("aaa")

	t.Log("Given the need to test the completion feature")
	{
		if !reflect.DeepEqual(ac1, words[:len(words)-1]) {
			t.Fatal("Should be able to autocomplete on a stem word by alphabetical order and then by length", ballotX)
		}

		ac2, _ := autoComplete.Complete("a")
		if !reflect.DeepEqual(ac2, words) {
			t.Fatal("Should be able to autocomplete on a stem word by alphabetical order and then by length", ballotX)
		}
		t.Log("Should be able to autocomplete on a stem word by alphabetical order and then by length", checkMark)

		_, err := NewAutoCompleteS(alphabet, []string{"漢", "字"}, 0, 0)
		if err == nil {
			t.Fatal("Should be able to reject non-alphabet words", ballotX)
		}
		t.Log("Should be able to reject non-alphabet words", checkMark)
	}
}

func Test_Learn(t *testing.T) {
	words := []string{"aaa", "b"}
	autoComplete, _ := NewAutoCompleteS(alphabet, words, 0, 0)
	err := autoComplete.Learn("aaabbb")
	if err != nil {
		t.Fatal(err)
	}

	ac, _ := autoComplete.Complete("aaa")
	t.Log("Given the need to test the Learn feature")
	{
		if !reflect.DeepEqual(ac, []string{"aaa", "aaabbb"}) {
			t.Fatal("Should be able to learn a new leaf word", ballotX)
		}
		t.Log("Should be able to learn a new leaf word", checkMark)
		if len(autoComplete.newWords) != 1 {
			t.Fatal("Should be able to correctly handle new word map", ballotX)
		}
		t.Log("Should be able to correctly handle new word map", checkMark)
		err := autoComplete.Learn("aa")
		if err != nil {
			t.Fatal(err)
		}
		ac, _ = autoComplete.Complete("aa")

		if !reflect.DeepEqual(ac, []string{"aa", "aaa", "aaabbb"}) {
			t.Fatal("Should be able to learn a new non-leaf word", ballotX)
		}
		t.Log("Should be able to learn a new non-leaf word", checkMark)
	}
	t.Log("Given the need to test the learn from scratch feature")
	{
		alphabet := "abcdefghijklmnopqrstuvwxyz"
		autoComplete, _ := NewAutoCompleteE(alphabet, 0, 0)

		ac, _ := autoComplete.Complete("aaa")
		if !reflect.DeepEqual(ac, []string{}) {
			t.Fatal("Should be able to correctly initialize an empty autocompleter", ballotX)
		}
		t.Log("Should be able to correctly initialize an empty autocompleter", checkMark)

		autoComplete.Learn("aaa")
		ac, _ = autoComplete.Complete("aaa")
		if !reflect.DeepEqual(ac, []string{"aaa"}) {
			t.Fatal("Should be able to learn from scratch", ballotX)
		}
		t.Log("Should be able to learn from scratch", checkMark)
	}
	t.Log("Given the need to test the UnLearn feature")
	{
		words := []string{"aaa", "aaab", "aaabbb", "aaabbbc", "ddd"}
		autoComplete, _ := NewAutoCompleteS(alphabet, words, 0, 0)
		autoComplete.UnLearn("aaabbbc")
		ac, _ := autoComplete.Complete("aaa")
		if !reflect.DeepEqual(ac, []string{"aaa", "aaab", "aaabbb"}) {
			t.Fatal("Should be able to unlearn a leaf", ballotX)
		}
		t.Log("Should be able to unlearn a leaf", checkMark)
		if len(autoComplete.newWords) != 0 {
			t.Fatal("Should be able to correctly handle new word map")
		}
		t.Log("Should be able to correctly handle new word map")
		autoComplete.UnLearn("aaabbb")
		autoComplete.UnLearn("aaab")
		ac, _ = autoComplete.Complete("aaa")
		if !reflect.DeepEqual(ac, []string{"aaa"}) {
			t.Fatal("Should be able to unlearn non-leaves", ballotX)
		}
		t.Log("Should be able to unlearn non-leaves", checkMark)
		autoComplete.UnLearn("aaa")
		ac, _ = autoComplete.Complete("ddd")
		if !reflect.DeepEqual(ac, []string{"ddd"}) {
			t.Fatal("Should be able to unlearn a branch", ballotX)
		}
		t.Log("Should be able to unlearn a branch", checkMark)
		autoComplete.UnLearn("ddd")
		ac, _ = autoComplete.Complete("ddd")
		for _, link := range autoComplete.root.links {
			if link != nil {
				t.Fatal("Should be able to unlearn whole tree", ballotX)
			}
		}
		t.Log("Should be able to unlearn whole tree", checkMark)
		removed := []string{}

		for w, _ := range autoComplete.removedWords {
			removed = append(removed, w)
		}
		sort.Strings(removed)
		sort.Strings(words)
		if !reflect.DeepEqual(removed, words) {
			t.Fatal("Should be able to correctly manage the removed word list", ballotX)
		}
		t.Log("Should be able to correctly manage the removed word list", checkMark)
		autoComplete.Learn("whatever")
		autoComplete.UnLearn("whatever")
		if !reflect.DeepEqual(removed, words) {
			t.Fatal("Should be able to correctly manage the removed word list for new words", ballotX)
		}
		t.Log("Should be able to correctly manage the removed word list for new words", checkMark)
	}
}
func Test_Accept(t *testing.T) {
	t.Log("Given the need to test the Accept feature")
	{
		words := []string{"aaa", "aaab", "aaac", "aaabbb", "aaad"} // Complete() always sorts by length and then alphabetically
		autoComplete, _ := NewAutoCompleteS(alphabet, words, 0, 0)
		autoComplete.Accept("aaad")
		ac, _ := autoComplete.Complete("aaa")

		if !reflect.DeepEqual(ac, []string{"aaad", "aaa", "aaab", "aaac", "aaabbb"}) {
			t.Fatal("Should be able to prioritize a word", ballotX)
		}
		t.Log("Should be able to prioritize a word", checkMark)
		autoComplete.Accept("aaad")
		autoComplete.Accept("aaac")
		ac, _ = autoComplete.Complete("aaa")
		if !reflect.DeepEqual(ac, []string{"aaad", "aaac", "aaa", "aaab", "aaabbb"}) {
			t.Fatal("Should be able to prioritize prioritized words", ballotX)
		}
		t.Log("Should be able to prioritize prioritized words", checkMark)
	}
}

func Test_ResultSizeAndRadius(t *testing.T) {
	t.Log("Given the need to test the Result size feature")
	{
		words := []string{"aaa", "aaab", "aaac", "aaabbb", "aaad"}
		autoComplete, _ := NewAutoCompleteS(alphabet, words, 3, 4)
		ac, _ := autoComplete.Complete("aaa")

		if !reflect.DeepEqual(ac, []string{"aaa", "aaab", "aaac"}) {
			t.Fatal("Should be able to limit result set size", ballotX)
		}
		t.Log("Should be able to limit result set size", checkMark)
	}
	t.Log("Given the need to test the radius feature")
	{
		numAlphabet := "1234567890"
		words := []string{"1234", "12345", "123456", "1234567", "12345678"}
		autoComplete, _ := NewAutoCompleteS(numAlphabet, words, 10, 4)
		ac, _ := autoComplete.Complete("1234")
		if !reflect.DeepEqual(ac, []string{"1234"}) {
			t.Fatal("Should be able to limit radius", ballotX)
		}
		t.Log("Should be able to limit radius", checkMark)
	}
}

func Test_SaveRetrieve(t *testing.T) {

	tempDir := os.TempDir()
	tempFile, err := ioutil.TempFile(tempDir, "smac")
	if err != nil {
		t.Fatal(err)
	}
	fName := tempFile.Name()
	words := []string{"aaa", "aaabbb", "bbb", "ccc"}
	autoComplete, _ := NewAutoCompleteS(alphabet, words, 0, 0)
	autoComplete.Accept("aaabbb")
	autoComplete.Learn("ddd")
	autoComplete.Learn("eee")
	autoComplete.Accept("eee")
	autoComplete.UnLearn("ccc")

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

		var wA wordAccepts
		dec.Decode(&wA)

		result1 := wordAccepts{
			"ddd",
			0,
		}
		if !reflect.DeepEqual(wA, result1) {
			t.Fatal("Should be able to read back a saved word", ballotX)
		}
		t.Log("Should be able to read back a saved word", checkMark)

		result2 := wordAccepts{
			"eee",
			1,
		}
		var wA2 wordAccepts
		dec.Decode(&wA2)
		if !reflect.DeepEqual(wA2, result2) {
			t.Fatal("Should be able to read back a saved and accepted word", ballotX)
		}
		t.Log("Should be able to read back a saved and accepted word", checkMark)

		result3 := wordAccepts{
			"aaabbb",
			1,
		}
		var wA3 wordAccepts
		dec.Decode(&wA3)
		if !reflect.DeepEqual(wA3, result3) {
			t.Fatal("Should be able to read back a second saved word", ballotX)
		}
		t.Log("Should be able to read back a second saved word", checkMark)

		autoComplete, _ = NewAutoCompleteS(alphabet, words, 0, 0)
		err = autoComplete.Retrieve(fName)
		if err != nil {
			t.Fatal(err)
		}
		ac, _ := autoComplete.Complete("aaa")
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
	}
}

func Example_test() {

	myAlphabet := "abcdefghijklmnopqrstuvwxyz"
	words := []string{"chair", "chairman", "chairperson", "chairwoman"}
	autoComplete, err := NewAutoCompleteS(myAlphabet, words, 0, 0)
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
