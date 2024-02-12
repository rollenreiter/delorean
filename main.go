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
	input := urls{
		source:  "",
		tokens:  []string{""},
		urls:    []string{""},
		results: []string{""},
	}
	flags := GetFlags()
	input.GetTokens(flags)
	fmt.Println(input)
	input.GetUrls(flags)
	fmt.Println(input)
	// input.Archive(flags)

	if flags.silentFlag {
		for i := range input.results {
			fmt.Printf("%s\n", input.results[i])
		}
	} else {
		fmt.Printf("These are the links to the archives:\n")
		for i := range input.results {
			fmt.Printf("%s\n", input.results[i])
		}
	}
}
