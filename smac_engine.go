package smac

import (
	"bufio"
	"errors"
	"os"
)

type trieNode struct {
	isWord  bool
	intRune int
	links   []*trieNode
}

type AutoComplete struct {
	root         *trieNode
	alphabetMin  int
	alphabetMax  int
	alphabetSize int
}

func NewAutoCompleteS(words []string) (AutoComplete, error) {
	var autoComplete AutoComplete

	min := 0
	max := 0
	for _, w := range words {
		runes := []rune(w)
		for _, r := range runes {
			if min > int(r) {
				min = int(r)
			}
			if max < int(r) {
				max = int(r)
			}
		}

	}
	autoComplete = AutoComplete{
		alphabetMin:  min,
		alphabetMax:  max,
		alphabetSize: max - min + 1,
	}

	autoComplete.root = &trieNode{
		links: make([]*trieNode, autoComplete.alphabetSize),
	}

	for _, w := range words {
		autoComplete.put(w)
	}

	return autoComplete, nil
}

func NewAutoCompleteF(fileName string) (AutoComplete, error) {

	var autoComplete AutoComplete

	f, err := os.Open(fileName)
	defer f.Close()

	if err != nil {
		return autoComplete, err
	}

	lineScanner := bufio.NewScanner(f)

	var min, max int

	for lineScanner.Scan() {
		line := lineScanner.Text()
		runes := []rune(line)
		for _, r := range runes {
			if min > int(r) {
				min = int(r)
			}
			if max < int(r) {
				max = int(r)
			}
		}
	}

	autoComplete = AutoComplete{
		alphabetMin:  min,
		alphabetMax:  max,
		alphabetSize: max - min + 1,
	}

	f.Seek(0, 0)
	lineScanner = bufio.NewScanner(f)

	autoComplete.root = &trieNode{
		links: make([]*trieNode, autoComplete.alphabetSize),
	}

	for lineScanner.Scan() {
		line := lineScanner.Text()
		autoComplete.put(line)
	}

	return autoComplete, nil
}

func (autoComplete *AutoComplete) runesToInts(word string) ([]int, error) {
	runes := []rune(word)
	var conv []int

	for _, r := range runes {
		i := int(r)
		if i < autoComplete.alphabetMin || i > autoComplete.alphabetMax {
			return nil, errors.New("illegal char in word - " + string(r))
		}
		conv = append(conv, i)
	}

	return conv, nil
}

func (autoComplete *AutoComplete) Learn(word string) error {
	conv, err := autoComplete.runesToInts(word)
	if err != nil {
		return err
	}
	autoComplete.putIter(conv)
	return nil
}

func (autoComplete *AutoComplete) put(word string) {

	conv, _ := autoComplete.runesToInts(word)
	autoComplete.putIter(conv)
}

func (autoComplete *AutoComplete) putIter(intVals []int) {

	node := autoComplete.root

	for i, c := range intVals {
		if node.links[c-autoComplete.alphabetMin] == nil {
			newNode := trieNode{
				intRune: c,
				links:   make([]*trieNode, autoComplete.alphabetSize),
			}
			if i == len(intVals)-1 {
				newNode.isWord = true
			}
			node.links[c-autoComplete.alphabetMin] = &newNode
			node = &newNode
			continue
		}
		node = node.links[c-autoComplete.alphabetMin]
		if i == len(intVals)-1 {
			node.isWord = true
		}
	}
}

func (autoComplete *AutoComplete) Complete(word string) ([]string, error) {

	ints, err := autoComplete.runesToInts(word)
	if err != nil {
		return nil, err
	}
	return autoComplete.complete(word, ints), nil
}

func (autoComplete *AutoComplete) complete(word string, intRunes []int) []string {

	wordEnd := autoComplete.root

	for _, c := range intRunes {
		wordEnd = wordEnd.links[c-autoComplete.alphabetMin]
		if wordEnd == nil {
			return []string{word}
		}
	}

	words := []string{}
	fifo := fIFO{}
	stem := []rune(word)
	stem = stem[:len(stem)-1]
	fifo.add(branch{
		node:   wordEnd,
		parent: stem,
	})

	for fifo.size() > 0 {
		nodeBranch := fifo.remove()
		if nodeBranch.node.isWord {
			words = append(words, string(append(nodeBranch.parent, rune(nodeBranch.node.intRune))))
		}
		links := nodeBranch.node.links

		for _, link := range links {
			parentString := make([]rune, len(nodeBranch.parent)+1)
			copy(parentString, nodeBranch.parent)
			parentString = append(parentString, rune(nodeBranch.node.intRune))
			if link != nil {
				rightBranch := branch{
					node:   link,
					parent: parentString,
				}
				fifo.add(rightBranch)
			}
		}
	}
	return words
}

type branch struct {
	node   *trieNode
	parent []rune
}

type fIFO struct {
	slice []branch
}

func (fifo *fIFO) add(b branch) {
	fifo.slice = append(fifo.slice, b)
}
func (fifo *fIFO) remove() branch {
	b := fifo.slice[0]
	fifo.slice = fifo.slice[1:]
	return b
}
func (fifo *fIFO) size() int {
	return len(fifo.slice)
}
