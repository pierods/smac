// Copyright Piero de Salvia.
// All Rights Reserved

package smac

import (
	"bufio"
	"encoding/gob"
	"errors"
	"io"
	"os"
)

type trieNode struct {
	isWord  bool
	intRune int
	accepts int
	links   []*trieNode
}

// AutoCompleteTrie represents the autocomplete engine.
type AutoCompleteTrie struct {
	root         *trieNode
	alphabetMin  int
	alphabetMax  int
	alphabetSize int
	resultSize   int
	radius       int
	newWords     map[string]byte
	removedWords map[string]byte
}

// NewAutoCompleteTrieE returns a new, empty autocompleter for a given alphabet (set of runes).
//
// resultSize is the number of hits returned. If 0 is used, it defaults to DEF_RESULTS_SIZE
//
// radius is the max length of words the engine will search while autocompleting. If 0 is used, it defaults to DEF_RADIUS
//
// The returned completer does not contain any words to be completed. New words can be added to it by using the Learn()
// function
func NewAutoCompleteTrieE(alphabet string, resultSize, radius uint) (AutoCompleteTrie, error) {

	var nAc AutoCompleteTrie
	if len(alphabet) == 0 {
		return nAc, errors.New("Empty alphabet")
	}
	min, max := minMax([]rune(alphabet))

	if resultSize == 0 {
		resultSize = DefaultResultSize
	}
	if radius == 0 {
		radius = DefaultRadius
	}

	autoComplete := AutoCompleteTrie{
		alphabetMin:  int(min),
		alphabetMax:  int(max),
		alphabetSize: int(max) - int(min) + 1,
		resultSize:   int(resultSize),
		radius:       int(radius),
		newWords:     make(map[string]byte),
		removedWords: make(map[string]byte),
	}

	autoComplete.root = &trieNode{
		links: make([]*trieNode, autoComplete.alphabetSize),
	}
	return autoComplete, nil
}

// NewAutoCompleteTrieS returns a new autocompleter for a given alphabet (set of runes).
//
// dictionary is a slice of words to be used for completion.
//
// resultSize is the number of hits returned. If 0 is used, it defaults to DEF_RESULTS_SIZE
//
// radius is the max length of words the engine will search while autocompleting. If 0 is used, it defaults to DEF_RADIUS
//
// New words can be added to it by using the Learn() function
func NewAutoCompleteTrieS(alphabet string, dictionary []string, resultSize, radius uint) (AutoCompleteTrie, error) {

	var nAc AutoCompleteTrie

	if len(alphabet) == 0 {
		return nAc, errors.New("Empty alphabet")
	}
	min, max := minMax([]rune(alphabet))

	if resultSize == 0 {
		resultSize = DefaultResultSize
	}
	if radius == 0 {
		radius = DefaultRadius
	}
	autoComplete := AutoCompleteTrie{
		alphabetMin:  int(min),
		alphabetMax:  int(max),
		alphabetSize: int(max) - int(min) + 1,
		resultSize:   int(resultSize),
		radius:       int(radius),
		newWords:     make(map[string]byte),
		removedWords: make(map[string]byte),
	}

	autoComplete.root = &trieNode{
		links: make([]*trieNode, autoComplete.alphabetSize),
	}

	for _, w := range dictionary {
		err := autoComplete.put(w)
		if err != nil {
			return nAc, err
		}
	}

	return autoComplete, nil
}

// NewAutoCompleteTrieF returns a new autocompleter for a given alphabet (set of runes).
//
// dictionaryFileName is the name of a dictionary file (a file containing words) to be used for completion.
//
// resultSize is the number of hits returned. If 0 is used, it defaults to DEF_RESULTS_SIZE
//
// radius is the max length of words the engine will search while autocompleting. If 0 is used, it defaults to DEF_RADIUS
//
// New words can be added to it by using the Learn() function
func NewAutoCompleteTrieF(alphabet, dictionaryFileName string, resultSize, radius uint) (AutoCompleteTrie, error) {

	var nAc AutoCompleteTrie
	if len(alphabet) == 0 {
		return nAc, errors.New("Empty alphabet")
	}
	min, max := minMax([]rune(alphabet))

	if resultSize == 0 {
		resultSize = DefaultResultSize
	}
	if radius == 0 {
		radius = DefaultRadius
	}

	autoComplete := AutoCompleteTrie{
		alphabetMin:  int(min),
		alphabetMax:  int(max),
		alphabetSize: int(max) - int(min) + 1,
		resultSize:   int(resultSize),
		radius:       int(radius),
		newWords:     make(map[string]byte),
		removedWords: make(map[string]byte),
	}

	f, err := os.Open(dictionaryFileName)
	defer f.Close()

	if err != nil {
		return nAc, err
	}

	lineScanner := bufio.NewScanner(f)

	autoComplete.root = &trieNode{
		links: make([]*trieNode, autoComplete.alphabetSize),
	}

	for lineScanner.Scan() {
		line := lineScanner.Text()
		err := autoComplete.put(line)
		if err != nil {
			return nAc, err
		}
	}

	return autoComplete, nil
}

// Accept : See description in AutoComplete interface
func (autoComplete *AutoCompleteTrie) Accept(acceptedWord string) error {
	acceptedWordInts, err := autoComplete.runesToInts(acceptedWord)
	if err != nil {
		return err
	}
	node := autoComplete.root
	for _, c := range acceptedWordInts {
		if node.links[c-autoComplete.alphabetMin] == nil {
			return errors.New("Word " + acceptedWord + " not in dictionary")
		}
		node = node.links[c-autoComplete.alphabetMin]
	}
	node.accepts++
	return nil
}

func (autoComplete *AutoCompleteTrie) runesToInts(word string) ([]int, error) {
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

// Learn : see interface
func (autoComplete *AutoCompleteTrie) Learn(word string) error {
	conv, err := autoComplete.runesToInts(word)
	if err != nil {
		return err
	}
	autoComplete.putIter(conv)
	autoComplete.newWords[word] = 0
	return nil
}

func (autoComplete *AutoCompleteTrie) put(word string) error {

	conv, err := autoComplete.runesToInts(word)
	if err != nil {
		return err
	}
	autoComplete.putIter(conv)
	return nil
}

func (autoComplete *AutoCompleteTrie) putIter(intVals []int) {

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

// UnLearn :  : See description in AutoComplete interface
func (autoComplete *AutoCompleteTrie) UnLearn(word string) error {
	conv, err := autoComplete.runesToInts(word)
	if err != nil {
		return err
	}
	autoComplete.remove(conv)
	if _, contains := autoComplete.newWords[word]; !contains {
		autoComplete.removedWords[word] = 0
	} else {
		delete(autoComplete.newWords, word)
	}

	return nil
}

func (autoComplete *AutoCompleteTrie) remove(intVals []int) {
	node := autoComplete.root
	lifo := lIFO{}

	for _, c := range intVals {
		if node.links[c-autoComplete.alphabetMin] == nil {
			return
		}
		lifo.push(node)
		node = node.links[c-autoComplete.alphabetMin]
	}
	if !node.isWord {
		return
	}
	isLeaf := true
	for _, link := range node.links {
		if link != nil {
			isLeaf = false
			break
		}
	}
	if !isLeaf {
		node.isWord = false
		return

	}
	node.isWord = false

	for lifo.size() > 0 {
		parentNode := lifo.pop()
		parentNode.links[node.intRune-autoComplete.alphabetMin] = nil

		if parentNode.isWord {
			return
		}
		for _, link := range parentNode.links {
			if link != nil {
				return
			}
		}
		node = parentNode
	}
}

// Complete : see description in Autocomplete interface
func (autoComplete *AutoCompleteTrie) Complete(word string) ([]string, error) {

	ints, err := autoComplete.runesToInts(word)
	if err != nil {
		return nil, err
	}
	return autoComplete.complete(word, ints), nil
}

func (autoComplete *AutoCompleteTrie) complete(word string, intRunes []int) []string {

	wordEnd := autoComplete.root
	for _, c := range intRunes {
		wordEnd = wordEnd.links[c-autoComplete.alphabetMin]
		if wordEnd == nil {
			return []string{}
		}
	}

	words := sOLILI{}
	fifo := fIFO{}
	stem := []rune(word)
	stem = stem[:len(stem)-1]
	fifo.add(branch{
		node:   wordEnd,
		parent: &stem,
	})
	results := 0
	for fifo.size() > 0 {
		if results == autoComplete.resultSize {
			break
		}

		nodeBranch := fifo.remove()
		if nodeBranch.node.isWord {
			words.insert(string(append(*nodeBranch.parent, rune(nodeBranch.node.intRune))), nodeBranch.node.accepts)
			results++
		}
		links := nodeBranch.node.links

		if len(*nodeBranch.parent) < autoComplete.radius-1 {
			parentString := make([]rune, len(*nodeBranch.parent)+1)
			copy(parentString, *nodeBranch.parent)
			parentString[len(parentString)-1] = rune(nodeBranch.node.intRune)

			for _, link := range links {
				if link != nil {
					rightBranch := branch{
						node:   link,
						parent: &parentString,
					}
					fifo.add(rightBranch)
				}
			}
		}
	}
	return words.flush()
}

type branch struct {
	node   *trieNode
	parent *[]rune
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

type lIFO struct {
	slice []*trieNode
}

func (lifo *lIFO) push(n *trieNode) {
	lifo.slice = append(lifo.slice, n)
}
func (lifo *lIFO) pop() *trieNode {
	node := lifo.slice[len(lifo.slice)-1]
	lifo.slice = lifo.slice[:len(lifo.slice)-1]
	return node
}
func (lifo *lIFO) size() int {
	return len(lifo.slice)
}

// Save : see description in AutoComplete interface
func (autoComplete *AutoCompleteTrie) Save(fileName string) error {

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	enc := gob.NewEncoder(f)

	fifo := fIFO{}
	var nSlice []rune

	fifo.add(branch{
		node:   autoComplete.root,
		parent: &nSlice,
	})

	for fifo.size() > 0 {

		nodeBranch := fifo.remove()
		if nodeBranch.node.isWord {
			currWord := string(append(*nodeBranch.parent, rune(nodeBranch.node.intRune)))
			if nodeBranch.node.accepts > 0 {
				enc.Encode(wordAccepts{
					currWord,
					nodeBranch.node.accepts,
				})

			} else if _, exists := autoComplete.newWords[currWord]; exists {
				enc.Encode(wordAccepts{
					currWord,
					nodeBranch.node.accepts,
				})
			}
		}
		links := nodeBranch.node.links

		var parentString []rune

		if *nodeBranch.parent != nil {
			parentString = make([]rune, len(*nodeBranch.parent)+1)
			copy(parentString, *nodeBranch.parent)
			parentString[len(parentString)-1] = rune(nodeBranch.node.intRune)

		} else {
			parentString = []rune{}
		}

		for _, link := range links {
			if link != nil {
				rightBranch := branch{
					node:   link,
					parent: &parentString,
				}
				fifo.add(rightBranch)
			}
		}
	}
	for w := range autoComplete.removedWords {
		enc.Encode(wordAccepts{
			w,
			-1,
		})
	}
	return f.Close()
}

// Retrieve : see description in AutoComplete interface
func (autoComplete *AutoCompleteTrie) Retrieve(fileName string) error {

	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(f)
	for {
		var wA wordAccepts
		if err = dec.Decode(&wA); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		runesAsInts, err := autoComplete.runesToInts(wA.Word)
		if err != nil {
			return err
		}
		node := autoComplete.root
		for _, c := range runesAsInts {
			if node.links[c-autoComplete.alphabetMin] == nil {
				err = autoComplete.Learn(wA.Word)
				if err != nil {
					return err
				}
			}
			node = node.links[c-autoComplete.alphabetMin]
		}
		if wA.Accepts > 0 {
			if err = autoComplete.updateAccepts(runesAsInts, wA.Accepts); err != nil {
				return err
			}
		} else if wA.Accepts < 0 {
			autoComplete.UnLearn(wA.Word)
		}
	}

	return nil
}

func (autoComplete *AutoCompleteTrie) updateAccepts(word []int, accepts int) error {

	node := autoComplete.root

	for _, c := range word {
		if node.links[c-autoComplete.alphabetMin] == nil {
			return errors.New("Word not found")
		}
		node = node.links[c-autoComplete.alphabetMin]
	}
	node.accepts = accepts
	return nil
}

func minMax(runes []rune) (rune, rune) {

	if len(runes) == 1 {
		return runes[0], runes[0]
	}

	var min, max rune

	if runes[0] > runes[1] {
		max = runes[0]
		min = runes[1]
	} else {
		max = runes[1]
		min = runes[0]
	}

	for i := 2; i < len(runes); i++ {
		if runes[i] > max {
			max = runes[i]
		} else if runes[i] < min {
			min = runes[i]
		}
	}

	return min, max
}
