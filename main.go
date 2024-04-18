package main

import (
	"fmt"
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

const Version string = "v1.1.0"

var Flags CmdArgs

func main() {
	Flags.SetFlags(FromStdin)
	if Flags.versionFlag {
		fmt.Printf("DeLorean-%s\nCopyright (C) 2024 Rollenreiter\nThis software may be freely redistributed under the terms of the GNU General Public License.\n", Version)
		os.Exit(0)
	}
	var wg sync.WaitGroup

	CheckConnection(&wg)

	input := NewInput()
	input.Tokenize()
	input.GetUrls(&wg)
	if len(input.validUrls) != 0 {
		input.Archive(&wg)
		Sort(input.results, Flags.alphaFlag)
		for _, s := range input.results {
			fmt.Println(s.content)
		}
	} else {
		if !Flags.silentFlag {

			fmt.Printf("%s|x|%s ", Error, Escape)
			fmt.Println("Nothing to archive.")
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
