// Copyright Piero de Salvia.
// All Rights Reserved

package smac

import (
	"bufio"
	"encoding/gob"
	"errors"
	"io"
	"os"
	"sort"
	"strings"
)

type liNo struct {
	accepts int
	next    string
}

type AutoCompleteLiNo struct {
	wordMap        map[string]*liNo
	head           string
	tail           string
	resultSize     int
	radius         int
	removedWords   map[string]bool
	newWords       map[string]bool
	prefixMap      map[string]string
	prefixMapDepth int
}

func NewAutoCompleteLinoE(prefixMapDepth, resultSize, radius uint) (AutoCompleteLiNo, error) {
	return NewAutoCompleteLinoS([]string{}, prefixMapDepth, resultSize, radius)
}

func NewAutoCompleteLinoF(dictionaryFileName string, prefixMapDepth, resultSize, radius uint) (AutoCompleteLiNo, error) {

	var nAc AutoCompleteLiNo

	f, err := os.Open(dictionaryFileName)
	defer f.Close()

	if err != nil {
		return nAc, err
	}

	lineScanner := bufio.NewScanner(f)
	var dictionary []string

	for lineScanner.Scan() {
		word := lineScanner.Text()
		if len(word) == 0 {
			return nAc, errors.New("Empty word in dictionary")
		}
		dictionary = append(dictionary, word)
	}

	return NewAutoCompleteLinoS(dictionary, prefixMapDepth, resultSize, radius)
}

func makePrefixMap(sortedDictionary []string, maxDepth int) map[string]string {

	prefixes := make(map[string]string)

	for _, word := range sortedDictionary {
		runes := []rune(word)
		for depth := 1; depth <= maxDepth && depth <= len(runes); depth++ {
			prefix := string(runes[:depth])
			if _, exists := prefixes[prefix]; !exists {
				prefixes[prefix] = word
			}
		}
	}
	return prefixes
}

func NewAutoCompleteLinoS(dictionary []string, prefixMapDepth, resultSize, radius uint) (AutoCompleteLiNo, error) {

	var nAc AutoCompleteLiNo

	if resultSize > radius {
		return nAc, errors.New("resultSize > radius")
	}

	if resultSize == 0 {
		resultSize = DEF_RESULTS_SIZE
	}
	if radius == 0 {
		radius = DEF_RADIUS
	}

	autoComplete := AutoCompleteLiNo{
		wordMap:      make(map[string]*liNo),
		resultSize:   int(resultSize),
		radius:       int(radius),
		newWords:     make(map[string]bool),
		removedWords: make(map[string]bool),
	}

	sort.Strings(dictionary)
	var linop *liNo

	for _, word := range dictionary {
		newLinop := &liNo{}
		autoComplete.wordMap[word] = newLinop
		if linop != nil {
			linop.next = word
		}
		linop = newLinop

	}
	autoComplete.head = dictionary[0]
	autoComplete.tail = dictionary[len(dictionary)-1]
	autoComplete.prefixMap = makePrefixMap(dictionary, int(prefixMapDepth))
	autoComplete.prefixMapDepth = int(prefixMapDepth)

	return autoComplete, nil
}

func (autoComplete *AutoCompleteLiNo) Complete(stem string) ([]string, error) {

	result := sOLILI{}
	hits := 0
	//radius := 0
	lino, hit := autoComplete.wordMap[stem]

	if hit {
		hits++
		result.insert(stem, lino.accepts)
	} else {
		subStem := stem
		prefixRoot, prefixExists := autoComplete.prefixMap[subStem]

		for !prefixExists && len(subStem) > 0 {
			subStem = subStem[:len(subStem)-1]
			prefixRoot, prefixExists = autoComplete.prefixMap[subStem]
		}
		if prefixExists {
			searchPtr := prefixRoot
			for !strings.HasPrefix(searchPtr, stem) {
				searchPtr = autoComplete.wordMap[searchPtr].next
				if searchPtr == "" || !strings.HasPrefix(searchPtr, subStem) {
					return []string{}, nil
				}
			}
			hit = true
			lino = autoComplete.wordMap[searchPtr]
			if _, prefixRootIsWord := autoComplete.wordMap[searchPtr]; prefixRootIsWord {
				hits++
				result.insert(searchPtr, lino.accepts)
			}
		}
	}
	for hit && hits < autoComplete.radius {
		word := lino.next
		hit = strings.HasPrefix(word, stem)
		if hit {
			lino = autoComplete.wordMap[word]
			hits++
			result.insert(word, lino.accepts)
		}
	}
	return result.flushL(autoComplete.resultSize), nil
}

func (autoComplete *AutoCompleteLiNo) Accept(acceptedWord string) error {

	var lino *liNo
	var exist bool
	if lino, exist = autoComplete.wordMap[acceptedWord]; !exist {
		return errors.New("Word to be accepted not found")
	}
	lino.accepts++
	return nil
}

func (autoComplete *AutoCompleteLiNo) Learn(word string) error {

	if _, exists := autoComplete.wordMap[word]; exists {
		return errors.New("Word already in dictionary")
	}

	prevWord := autoComplete.findPreviousWord(word)
	prevLino := autoComplete.wordMap[prevWord]

	newLino := &liNo{}
	autoComplete.wordMap[word] = newLino

	if prevWord != autoComplete.head {
		newLino.next = prevLino.next
		prevLino.next = word
	} else {
		newLino.next = autoComplete.head
	}

	if autoComplete.head > word {
		autoComplete.head = word
	}
	if autoComplete.tail < word {
		autoComplete.tail = word
	}

	for i := 1; i <= autoComplete.prefixMapDepth && i <= len(word); i++ {
		prefix := word[:i]
		if _, exists := autoComplete.prefixMap[prefix]; !exists {
			autoComplete.prefixMap[prefix] = word
		} else {
			if word < autoComplete.prefixMap[prefix] {
				autoComplete.prefixMap[prefix] = word
			}
		}
	}
	autoComplete.newWords[word] = true
	return nil
}

func (autoComplete *AutoCompleteLiNo) UnLearn(word string) error {

	if _, exists := autoComplete.wordMap[word]; !exists {
		return errors.New("Word not in dictionary")
	}

	prevWord := autoComplete.findPreviousWord(word)
	var nextWord string
	prevLino := autoComplete.wordMap[prevWord]
	if autoComplete.wordMap[word].next == "" {
		prevLino.next = ""
	} else {
		nextWord = autoComplete.wordMap[word].next
		prevLino.next = nextWord
	}

	delete(autoComplete.wordMap, word)

	for i := 1; i <= autoComplete.prefixMapDepth && i < len(word); i++ {
		prefix := word[:i]
		if _, exists := autoComplete.prefixMap[prefix]; exists {
			if autoComplete.prefixMap[prefix] == word {
				// does next word start with prefix? if yes, assign, otherwise prefix is gone
				if strings.HasPrefix(nextWord, prefix) {
					autoComplete.prefixMap[prefix] = nextWord
				} else {
					delete(autoComplete.prefixMap, prefix)
				}
			}
		}
	}
	if autoComplete.head == word {
		autoComplete.head = nextWord
	}
	if autoComplete.tail == word {
		autoComplete.tail = prevWord
	}

	if _, contains := autoComplete.newWords[word]; !contains {
		autoComplete.removedWords[word] = true
	} else {
		delete(autoComplete.newWords, word)
	}

	return nil
}

func (autoComplete *AutoCompleteLiNo) findPreviousWord(word string) string {

	prefix := word[:len(word)-1]
	searchPtr, prefixExists := autoComplete.prefixMap[prefix]

	for len(prefix) > 0 && (!prefixExists || word <= searchPtr) {
		prefix = prefix[:len(prefix)-1]
		searchPtr, prefixExists = autoComplete.prefixMap[prefix]
	}
	// find the longest prefix present in prefixMap
	if searchPtr == "" { // prefix not found
		searchPtr = autoComplete.head
	}
	prevWord := searchPtr
	// now scan from longest prefix ptr until next word in dictionary is found
	for searchPtr != "" && searchPtr < word {
		prevWord = searchPtr
		searchPtr = autoComplete.wordMap[searchPtr].next
	}

	return prevWord
}

func (autoComplete *AutoCompleteLiNo) Save(fileName string) error {

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	enc := gob.NewEncoder(f)

	for w, liNo := range autoComplete.wordMap {
		if liNo.accepts > 0 {
			enc.Encode(wordAccepts{
				w,
				liNo.accepts,
			})
		} else if _, exists := autoComplete.newWords[w]; exists {
			enc.Encode(wordAccepts{
				w,
				liNo.accepts,
			})
		}
	}

	for w, _ := range autoComplete.removedWords {
		enc.Encode(wordAccepts{
			w,
			-1,
		})
	}
	return f.Close()
}

func (autoComplete *AutoCompleteLiNo) Retrieve(fileName string) error {
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
		if _, exists := autoComplete.wordMap[wA.Word]; !exists {
			err = autoComplete.Learn(wA.Word)
			if err != nil {
				return err
			}
		}
		if wA.Accepts > 0 {
			l := autoComplete.wordMap[wA.Word]
			l.accepts = wA.Accepts
		} else if wA.Accepts < 0 {
			autoComplete.UnLearn(wA.Word)
		}
	}
	return nil
}
