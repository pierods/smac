package smac

import (
	"bufio"
	"errors"
	"os"
)

const DEF_RESULTS_SIZE = 10
const DEF_RADIUS = 15

type trieNode struct {
	isWord  bool
	intRune int
	accepts int
	links   []*trieNode
}

type AutoComplete struct {
	root         *trieNode
	alphabetMin  int
	alphabetMax  int
	alphabetSize int
	resultSize   int
	radius       int
}

func NewAutoCompleteE(alphabet string, resultSize, radius uint) (AutoComplete, error) {

	var min, max int

	for _, r := range alphabet {
		i := int(r)
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}

	if resultSize == 0 {
		resultSize = DEF_RESULTS_SIZE
	}
	if radius == 0 {
		radius = DEF_RADIUS
	}

	autoComplete := AutoComplete{
		alphabetMin:  min,
		alphabetMax:  max,
		alphabetSize: max - min + 1,
		resultSize:   int(resultSize),
		radius:       int(radius),
	}

	autoComplete.root = &trieNode{
		links: make([]*trieNode, autoComplete.alphabetSize),
	}
	return autoComplete, nil
}

func NewAutoCompleteS(words []string, resultSize, radius uint) (AutoComplete, error) {

	var min, max int

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

	if resultSize == 0 {
		resultSize = DEF_RESULTS_SIZE
	}
	if radius == 0 {
		radius = DEF_RADIUS
	}
	autoComplete := AutoComplete{
		alphabetMin:  min,
		alphabetMax:  max,
		alphabetSize: max - min + 1,
		resultSize:   int(resultSize),
		radius:       int(radius),
	}

	autoComplete.root = &trieNode{
		links: make([]*trieNode, autoComplete.alphabetSize),
	}

	for _, w := range words {
		autoComplete.put(w)
	}

	return autoComplete, nil
}

func NewAutoCompleteF(fileName string, resultSize, radius uint) (AutoComplete, error) {

	f, err := os.Open(fileName)
	defer f.Close()

	if err != nil {
		return AutoComplete{}, err
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

	if resultSize == 0 {
		resultSize = DEF_RESULTS_SIZE
	}
	if radius == 0 {
		radius = DEF_RADIUS
	}
	autoComplete := AutoComplete{
		alphabetMin:  min,
		alphabetMax:  max,
		alphabetSize: max - min + 1,
		resultSize:   int(resultSize),
		radius:       int(radius),
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

func (autoComplete *AutoComplete) Accept(acceptedWord string) {
	acceptedWordInts, err := autoComplete.runesToInts(acceptedWord)
	if err != nil {
		return
	}
	node := autoComplete.root
	for _, c := range acceptedWordInts {
		if node.links[c-autoComplete.alphabetMin] == nil {
			return
		}
		node = node.links[c-autoComplete.alphabetMin]
	}
	node.accepts++
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

func (autoComplete *AutoComplete) UnLearn(word string) error {
	conv, err := autoComplete.runesToInts(word)
	if err != nil {
		return err
	}
	autoComplete.remove(conv)
	return nil
}

func (autoComplete *AutoComplete) remove(intVals []int) {
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

	} else {
		node.isWord = false

		for lifo.size() > 0 {
			parentNode := lifo.pop()
			parentNode.links[node.intRune] = nil

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

		if len(*nodeBranch.parent) < autoComplete.radius {
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

type wordHit struct {
	word    string
	accepts int
	next    *wordHit
}

type sOLILI struct {
	start *wordHit
	end   *wordHit
}

func (list *sOLILI) insert(word string, accepts int) {
	hit := &wordHit{
		word:    word,
		accepts: accepts,
	}
	if list.start == nil {
		list.start = hit
		list.end = hit
		return
	}

	if accepts == 0 {
		list.end.next = hit
		list.end = hit
		return
	}

	if hit.accepts > list.start.accepts {
		hit.next = list.start
		list.start = hit
		return
	}
	cursor := list.start
	for cursor.next != nil {
		if hit.accepts > cursor.next.accepts {
			break
		}
		cursor = cursor.next
	}
	hit.next = cursor.next
	cursor.next = hit
}

func (list *sOLILI) flush() []string {

	slice := []string{}

	if list.start != nil {
		cursor := list.start
		for cursor != nil {
			slice = append(slice, cursor.word)
			cursor = cursor.next
		}
	}

	return slice
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
