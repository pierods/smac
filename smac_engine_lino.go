package smac

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strings"
)

type liNo struct {
	accepts int
	next    string
}

type AutoCompleteLiNo struct {
	wordMap      map[string]*liNo
	head         string
	tail         string
	resultSize   int
	radius       int
	removedWords []string
	prefixMap    map[string]string
}

func NewAutoCompleteLinoF(dictionaryFileName string, maxDepth, resultSize, radius uint) (AutoCompleteLiNo, error) {

	if resultSize == 0 {
		resultSize = DEF_RESULTS_SIZE
	}
	if radius == 0 {
		radius = DEF_RADIUS
	}

	var nAc AutoCompleteLiNo

	autoComplete := AutoCompleteLiNo{
		wordMap:    make(map[string]*liNo),
		resultSize: int(resultSize),
		radius:     int(radius),
	}

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
	autoComplete.prefixMap = makePrefixMap(dictionary, int(maxDepth))
	return autoComplete, nil
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

func NewAutoCompleteLinoS(dictionary []string, maxDepth, resultSize, radius uint) (AutoCompleteLiNo, error) {

	if resultSize == 0 {
		resultSize = DEF_RESULTS_SIZE
	}
	if radius == 0 {
		radius = DEF_RADIUS
	}

	autoComplete := AutoCompleteLiNo{
		wordMap:    make(map[string]*liNo),
		resultSize: int(resultSize),
		radius:     int(radius),
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
	autoComplete.prefixMap = makePrefixMap(dictionary, int(maxDepth))

	return autoComplete, nil
}

func (autoComplete *AutoCompleteLiNo) Complete(stem string) ([]string, error) {

	result := []string{}

	lino, hit := autoComplete.wordMap[stem]

	if hit {
		result = append(result, stem)
	}
	for hit {
		word := lino.next
		hit = strings.HasPrefix(word, stem)
		if hit {
			result = append(result, word)
			lino = autoComplete.wordMap[word]
		}
	}
	sort.Sort(byLen(result))
	return result, nil
}

type byLen []string

func (a byLen) Len() int           { return len(a) }
func (a byLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func (autoComplete *AutoCompleteLiNo) Accept(acceptedWord string) error {

	var lino *liNo
	var exist bool
	if lino, exist = autoComplete.wordMap[acceptedWord]; !exist {
		return errors.New("Word to be accepted not found")
	}

	lino.accepts++
	return nil
}

func (autoComplete *AutoCompleteLiNo) UnLearn(word string) error {

	return nil
}
