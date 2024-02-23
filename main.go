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

func main() {
	input := NewInput()
	flags := GetFlags()
	var wg sync.WaitGroup
	if !flags.silentFlag {
		input.Tokenize(flags)
		input.GetUrls(flags, &wg)
		fmt.Println("Done.")
		input.Archive(flags, &wg)
		fmt.Printf("\nSUCCESS! These are the links to the archives:\n")

		Sort(input.results, flags.alphaFlag)
		for _, s := range input.results {
			fmt.Println(s.content)
		}
	} else {
		input.TokenizeSilent(flags)
		input.GetUrlsSilent(flags, &wg)
		input.ArchiveSilent(flags, &wg)

		Sort(input.results, flags.alphaFlag)
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
