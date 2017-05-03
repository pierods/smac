package smac

import (
	"reflect"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func Test_fIFO(t *testing.T) {

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
			a := branch{
				node:   nil,
				parent: []rune("a"),
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
			b := branch{
				node:   nil,
				parent: []rune("b"),
			}
			c := branch{
				node:   nil,
				parent: []rune("c"),
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

func Test_RunesToInts(t *testing.T) {

	words := []string{"aaa", "aaaa", "aaab", "aaac"}

	autoComplete, _ := NewAutoCompleteS(words)

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

	autoComplete, _ := NewAutoCompleteS(words)

	t.Log("Given the need to test the putIter() function")
	{
		a := autoComplete.root.links[97]
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
		b := a.links[98]
		if b.intRune != 98 {
			t.Fatal("Should be able to insert second character of a word in tree", ballotX)
		}
		if b.isWord {
			t.Fatal("Should be able to insert second character of a word in tree", ballotX)
		}
		t.Log("Should be able to insert second character of a word in tree", checkMark)
		c := b.links[99]
		if !c.isWord {
			t.Fatal("Should be able to insert a word in tree", ballotX)
		}
		t.Log("Should be able to insert a word in tree", checkMark)
	}
}

func Test_Completion(t *testing.T) {

	words := []string{"aaa", "aaab", "aaac", "aaad", "abbbbb"}
	autoComplete, _ := NewAutoCompleteS(words)
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

	}
}
