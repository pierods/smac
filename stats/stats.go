package main

import (
	"bufio"
	"fmt"
	"os"
	//"sort"
	//"strings"
)

func main() {

	f, err := os.Open(os.Args[1])
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	lineScanner := bufio.NewScanner(f)
	var dictionary []string
	dictionaryMap := make(map[string]bool)

	for lineScanner.Scan() {
		word := lineScanner.Text()
		if len(word) == 0 {
			fmt.Println("Empty word in dictionary")
			os.Exit(-1)
		}
		dictionary = append(dictionary, word)
		dictionaryMap[word] = true
	}

	prefixes3 := make(map[string]int)
	prefixes2 := make(map[string]int)
	prefixes1 := make(map[string]int)
	fmt.Println("Building prefix counts...")
	for _, word := range dictionary {
		if len(word) > 2 {
			prefix3 := word[:3]
			prefixes3[prefix3]++
		}
		if len(word) > 1 {
			prefix2 := word[:2]
			prefix1 := word[:1]
			prefixes2[prefix2]++
			prefixes1[prefix1]++
		} else {
			prefixes1[word]++
		}
	}

	fmt.Println("Calculating mean # of words for all prefixes...")

	mean3 := 0
	words3 := 0

	for w, v := range prefixes3 {
		//fmt.Println(k, " ", v)
		mean3 += v
		if _, exists := dictionaryMap[w]; exists {
			words3++
		}
	}
	mean2 := 0
	for _, v := range prefixes2 {
		//fmt.Println(k, " ", v)
		mean2 += v
	}

	mean1 := 0
	for _, v := range prefixes1 {
		//fmt.Println(k, " ", v)
		mean1 += v
	}

	fmt.Println("mean # of words for all order 1 prefixes: ", mean1/len(prefixes1))
	fmt.Println("mean # of words for all order 2 prefixes: ", mean2/len(prefixes2))
	fmt.Println("mean # of words for all order 3 prefixes: ", mean3/len(prefixes3))
	fmt.Println("% of whole words in order 3 prefixes: ", float64(words3)/float64(len(prefixes3)))
}
