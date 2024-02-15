package main

import (
	"fmt"
	"sync"
)

type urls struct {
	source    string
	tokens    []string
	validUrls []string
	results   []string
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

		for i := range input.results {
			fmt.Println(input.results[i])
		}
	} else {
		input.TokenizeSilent(flags)
		input.GetUrlsSilent(flags, &wg)
		input.ArchiveSilent(flags)

		for i := range input.results {
			fmt.Println(input.results[i])
		}
	}
}

func NewInput() urls {
	i := urls{
		source:    "",
		tokens:    make([]string, 0),
		validUrls: make([]string, 0),
		results:   make([]string, 0),
	}
	return i
}
