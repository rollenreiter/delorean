package main

import (
	"fmt"
)

type urls struct {
	source  string
	tokens  []string
	urls    []string
	results []string
}

func main() {
	input := NewInput()
	flags := GetFlags()
	input.GetTokens(flags)
	input.GetUrls(flags)
	input.Archive(flags)

	if flags.silentFlag {
		for i := range input.results {
			fmt.Printf("%s\n", input.results[i])
		}
	} else {
		fmt.Printf("These are the links to the archives:\n")
		for i := range input.results {
			fmt.Println(input.results[i])
		}
	}
}

func NewInput() urls {
	i := urls{
		source:  "",
		tokens:  make([]string, 0),
		urls:    make([]string, 0),
		results: make([]string, 0),
	}
	return i
}
