package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
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
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		if err != nil {
			rw.Write([]byte(""))
			return
		}
		json.NewEncoder(rw).Encode(&completions)
		return
	}

	if strings.HasPrefix(path, "/accept") {
		splitPath := strings.Split(path, "/")
		acceptedWord := splitPath[2]

		err := autoComplete.Accept(acceptedWord)
		if err != nil {
			err = autoComplete.Learn(acceptedWord)
			if err != nil {
				autoComplete.Accept(acceptedWord)
			}
		}
		rw.WriteHeader(200)
		return
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

	f.Close()

	wordFile := goPath + "/src/github.com/pierods/smac/demo/allwords.txt"

	var err error

	autoCompleteL, err := smac.NewAutoCompleteLinoF(wordFile, 2, 0, 0)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	autoComplete = &autoCompleteL

	watcher := watch(homeFile)
	defer watcher.Close()
	http.HandleFunc("/", handler)
	fmt.Println("Listener : Started : Listening on port 30000")
	http.ListenAndServe(":30000", nil)
}

func watch(fileName string) *fsnotify.Watcher {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("re-reading file:", event.Name)
					f, fErr := os.Open(fileName)
					if fErr != nil {
						fmt.Println(fErr)
						os.Exit(-1)
					}
					home, fErr = ioutil.ReadAll(f)
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return watcher
}
