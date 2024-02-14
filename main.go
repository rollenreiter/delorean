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
	if !flags.silentFlag {
		// input.GetTokens(flags)
		// input.GetUrls(flags)
		// input.Archive(flags)
		var wg sync.WaitGroup
		input.GetTokens(flags)
		input.GetUrls(flags, &wg)
		fmt.Println("Done.")
		input.Archive(flags, &wg)
	} else {
		// TODO: Make silent functions async.
		input.GetTokensSilent(flags)
		input.GetUrlsSilent(flags)
		input.ArchiveSilent(flags)
	}

	if flags.silentFlag {
		for i := range input.results {
			fmt.Println(input.results[i])
		}
	} else {
		fmt.Printf("\nSUCCESS! These are the links to the archives:\n")
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
