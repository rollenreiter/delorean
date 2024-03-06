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

const Version string = "v1.0.0"

var Flags cmdflags

func main() {
	Flags.SetFlags(FromStdin)
	if Flags.versionFlag {
		fmt.Printf("DeLorean-%s\nCopyright (C) 2024 Rollenreiter\nThis software may be freely redistributed under the terms of the GNU General Public License\n", Version)
		os.Exit(0)
	}
	input := NewInput()
	var wg sync.WaitGroup

	if !Flags.silentFlag {
		fmt.Printf("Checking availability...\n")
	}
	_, err := http.Get("http://web.archive.org")
	if err != nil {
		fmt.Printf("Couldn't connect to the Internet Archive. Please check your internet connection.\n")
		os.Exit(1)
	}
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
