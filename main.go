package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
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

var Flags CmdArgs

func main() {
	Flags.SetFlags(FromStdin)
	if Flags.versionFlag {
		fmt.Printf("DeLorean-%s\nCopyright (C) 2024 Rollenreiter\nThis software may be freely redistributed under the terms of the GNU General Public License\n", Version)
		os.Exit(0)
	}

	if !Flags.silentFlag {
		fmt.Printf("Checking availability...\n")
	}
	var wg sync.WaitGroup
	wg.Add(1)
	netfailure := make(chan bool, 1)
	go func() {
		_, err := http.Get("http://web.archive.org")
		if err != nil {
			fmt.Println("Couldn't connect to the Internet Archive. Please check your internet connection.")
			netfailure <- true
		} else {
			wg.Done()
		}
	}()
	select {
	case <-netfailure:
		os.Exit(1)
	case <-time.After(8 * time.Second):
		fmt.Println("Timed out while trying to connect to the Internet Archive. Please check your internet connection.")
		os.Exit(1)
	}
	wg.Wait()

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
