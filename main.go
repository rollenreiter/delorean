package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

type urls struct {
	source    string
	tokens    []token
	validUrls []token
	results   []token
}

type token struct {
	content string
	order   int
}

var Flags cmdflags

func main() {
	Flags.SetFlags(FromStdin)
	input := NewInput()
	var wg sync.WaitGroup

	fmt.Printf("Checking availability...\n")
	_, err := http.Get("http://web.archive.org")
	if err != nil {
		fmt.Printf("Couldn't connect to the Internet Archive. Please check your internet connection.\n")
		os.Exit(1)
	}
	input.Tokenize()
	input.GetUrls(&wg)
	input.Archive(&wg)
	Sort(input.results, Flags.alphaFlag)
	for _, s := range input.results {
		fmt.Println(s.content)
	}
}

func NewInput() urls {
	i := urls{
		source:    "",
		tokens:    make([]token, 0),
		validUrls: make([]token, 0),
		results:   make([]token, 0),
	}
	return i
}
