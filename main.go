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
