package main

import (
	"fmt"
	"sync"
)

type urls struct {
	source    string
	tokens    []token
	validUrls []token
	results   []token
}

type token struct {
	order   int
	content string
}

var Flags *cmdflags = GetFlags()

func main() {
	input := NewInput()
	var wg sync.WaitGroup
	if !Flags.silentFlag {
		input.Tokenize()
		input.GetUrls(&wg)
		fmt.Println("Done.")
		input.Archive(&wg)
		fmt.Printf("\nSUCCESS! These are the links to the archives:\n")

		Sort(input.results, Flags.alphaFlag)
		for _, s := range input.results {
			fmt.Println(s.content)
		}
	} else {
		input.TokenizeSilent()
		input.GetUrlsSilent(&wg)
		input.ArchiveSilent(&wg)

		Sort(input.results, Flags.alphaFlag)
		for _, s := range input.results {
			fmt.Println(s.content)
		}
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
