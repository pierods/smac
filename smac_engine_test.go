package smac

import (
	"reflect"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func Test_SOLILI(t *testing.T) {

	t.Log("Given the need to test a sorted linked list")
	{
		t.Log("When initializing a solili")
		{
			list := SOLILI{}
			slice := list.Flush()
			if !reflect.DeepEqual(slice, []string{}) {
				t.Fatal("Should be able to correctly initialize an empty solili", ballotX)
			}
			t.Log("Should be able to correctly initialize an empty solili", checkMark)
			list.Insert("aaa", 0)
			slice = list.Flush()
			if !reflect.DeepEqual(slice, []string{"aaa"}) {
				t.Fatal("Should be able to correctly add an element to a solili", ballotX)
			}
			t.Log("Should be able to correctly add an element to a solili", checkMark)
			list.Insert("bbb", 0)
			slice = list.Flush()
			if !reflect.DeepEqual(slice, []string{"aaa", "bbb"}) {
				t.Fatal("Should be able to correctly add an element to the back of a solili", ballotX)
			}
			t.Log("Should be able to correctly add an element to the back of a solili", checkMark)
			list.Insert("jjj", 100)
			slice = list.Flush()
			if !reflect.DeepEqual(slice, []string{"jjj", "aaa", "bbb"}) {
				t.Fatal("Should be able to correctly add an element to the front of a solili", ballotX)
			}
			t.Log("Should be able to correctly add an element to the front of a solili", checkMark)
			list.Insert("kkk", 90)
			slice = list.Flush()
			if !reflect.DeepEqual(slice, []string{"jjj", "kkk", "aaa", "bbb"}) {
				t.Fatal("Should be able to correctly add an element in the middle of a solili", ballotX)
			}
			t.Log("Should be able to correctly add an element in the middle of a solili", checkMark)
			list.Insert("lll", 100)
			slice = list.Flush()
			if !reflect.DeepEqual(slice, []string{"jjj", "lll", "kkk", "aaa", "bbb"}) {
				t.Fatal("Should be able to maintain arrival order for same-weight elements (at head)", ballotX)
			}
			t.Log("Should be able to maintain arrival order for same-weight elements (at head)", checkMark)
			list.Insert("mmm", 90)
			slice = list.Flush()
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

func Test_Learn(t *testing.T) {
	words := []string{"aaa", "b"}
	autoComplete, _ := NewAutoCompleteS(words)
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
		t.Log("Should be able to learn a new word", checkMark)

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
		autoComplete, _ := NewAutoCompleteE(alphabet)

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
		autoComplete, _ := NewAutoCompleteS(words)
		autoComplete.UnLearn("aaabbbc")
		ac, _ := autoComplete.Complete("aaa")
		if !reflect.DeepEqual(ac, []string{"aaa", "aaab", "aaabbb"}) {
			t.Fatal("Should be able to unlearn a leaf", ballotX)
		}
		t.Log("Should be able to unlearn a leaf", checkMark)
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
		for i, link := range autoComplete.root.links {
			if link != nil {
				t.Log(i)
				t.Fatal("Should be able to unlearn whole tree", ballotX)
			}
		}
		t.Log("Should be able to unlearn whole tree", checkMark)
	}
}
func Test_Accept(t *testing.T) {
	t.Log("Given the need to test the Accept feature")
	{
		words := []string{"aaa", "aaab", "aaac", "aaabbb", "aaad"} // Complete() always sorts by length and then alphabetically
		autoComplete, _ := NewAutoCompleteS(words)
		autoComplete.Accept("aaad")
		ac, _ := autoComplete.Complete("aaa")

		if !reflect.DeepEqual(ac, []string{"aaad", "aaa", "aaab", "aaac", "aaabbb"}) {
			t.Fatal("Should be able to prioritize a word", ballotX)
		}
		t.Log("Should be able to prioritize a word", checkMark)
		autoComplete.Accept("aaad")
		autoComplete.Accept("aaac")
		ac, _ = autoComplete.Complete("aaa")
		t.Log(ac)
		if !reflect.DeepEqual(ac, []string{"aaad", "aaac", "aaa", "aaab", "aaabbb"}) {
			t.Fatal("Should be able to prioritize prioritized words", ballotX)
		}
		t.Log("Should be able to prioritize prioritized words", checkMark)
	}
}