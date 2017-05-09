package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pierods/smac"
)

var (
	autoComplete smac.AutoComplete
	home         []byte
)

func handler(rw http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	if strings.HasPrefix(path, "/complete") {
		splitPath := strings.Split(path, "/")
		stem := splitPath[2]

		completions, err := autoComplete.Complete(stem)
		rw.Header().Set("Content/Type", "application/json")
		rw.WriteHeader(200)
		if err != nil {
			rw.Write([]byte(""))
			return
		}
		json.NewEncoder(rw).Encode(&completions)
		return
	}

	if strings.HasPrefix(path, "/accept") {

	}

	if strings.HasPrefix(path, "/learn") {

	}

	rw.Header().Set("Content/Type", "text/html")
	rw.WriteHeader(200)
	rw.Write(home)
}

func main() {

	goPath := os.Getenv("GOPATH")

	homeFile := goPath + "/src/github.com/pierods/smac/demo/demo.html"
	f, fErr := os.Open(homeFile)
	if fErr != nil {
		fmt.Println(fErr)
		os.Exit(-1)
	}
	home, fErr = ioutil.ReadAll(f)

	if fErr != nil {
		fmt.Println(fErr)
		os.Exit(-1)
	}

	wordFile := goPath + "/src/github.com/pierods/smac/demo/allwords.txt"
	alphabet := "abcdefghijklmnopqrstuvwxyz1234567890'/&\""

	var err error

	autoComplete, err = smac.NewAutoCompleteF(alphabet, wordFile, 0, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listener : Started : Listening on port 30000")
	http.ListenAndServe(":30000", nil)
}
