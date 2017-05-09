package main

import (
	"fmt"
	"os"

	"github.com/pierods/smac"
)

func main() {

	goPath := os.Getenv("GOPATH")
	wordFile := goPath + "/src/github.com/pierods/smac/demo/allwords.txt"
	alphabet := "abcdefghijklmnopqrstuvwxyz1234567890'/&\""
	autoComplete, err := smac.NewAutoCompleteF(alphabet, wordFile, 0, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	completes, err := autoComplete.Complete("chair")
	fmt.Println(completes)
}
