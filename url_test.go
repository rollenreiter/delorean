package main

import (
	"fmt"
	"sync"
	"testing"
)

// TODO: Adjust test cases for new async functions
func TestTokenizeUrl(t *testing.T) {
	flags := cmdflags{
		silentFlag: false,
		urlFlag:    "https://github.com/rollenreiter/delorean junktext http://asahina.moe github.com/stompiegit",
		fileFlag:   "",
	}
	input := NewInput()
	input.GetTokens(&flags)

	fmt.Printf("source is the following: %s\n", input.source)

	got := input.tokens
	fmt.Printf("got following tokens in input.tokens:\n%s\n", got)

	want := []string{"https://github.com/rollenreiter/delorean", "junktext", "http://asahina.moe", "github.com/stompiegit"}
	fmt.Printf("want following tokens:\n%s\n", want)

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestTokenizeFile(t *testing.T) {
	flags := cmdflags{
		silentFlag: false,
		urlFlag:    "",
		fileFlag:   "testfile",
	}
	input := NewInput()
	input.GetTokens(&flags)

	fmt.Printf("source is the following: %s\n", input.source)
	got := input.tokens
	fmt.Printf("got following tokens in input.tokens:\n%s\n", got)

	want := []string{
		"foo",
		"bar",
		"https://github.com/rollenreiter/delorean",
		"junktext",
		"http://asahina.moe",
		"github.com/stompiegit",
	}

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestGetUrls(t *testing.T) {
	flags := cmdflags{
		silentFlag: false,
		urlFlag:    "https://github.com/rollenreiter/delorean junktext http://asahina.moe github.com/stompiegit",
		fileFlag:   "",
	}
	var wg sync.WaitGroup
	input := NewInput()
	input.GetTokens(&flags)
	input.GetUrls(&flags, &wg)
	got := input.validUrls
	fmt.Println(input)
	want := []string{
		"https://github.com/rollenreiter/delorean",
		"http://asahina.moe",
		"http://github.com/stompiegit",
	}

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestGetUrlsFile(t *testing.T) {
	flags := cmdflags{
		silentFlag: false,
		urlFlag:    "",
		fileFlag:   "testfile",
	}
	var wg sync.WaitGroup
	input := NewInput()
	input.GetTokens(&flags)
	fmt.Println(input.tokens)
	input.GetUrls(&flags, &wg)
	got := input.validUrls
	fmt.Println(got)
	want := []string{
		"https://github.com/rollenreiter/delorean",
		"http://asahina.moe",
		"http://github.com/stompiegit",
	}

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
