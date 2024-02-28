package main

import (
	"fmt"
	"sync"
	"testing"
)

func FromTestURL() cmdflags {
	f := cmdflags{
		urlFlag:    "https://github.com/rollenreiter/delorean junktext http://asahina.moe github.com/stompman34",
		fileFlag:   "",
		silentFlag: false,
		alphaFlag:  false,
	}
	return f
}

func FromTestfile() cmdflags {
	f := cmdflags{
		urlFlag:    "",
		fileFlag:   "testfile",
		silentFlag: false,
		alphaFlag:  false,
	}
	return f
}

func TestTokenizeUrl(t *testing.T) {
	Flags.SetFlags(FromTestURL)
	input := NewInput()
	input.Tokenize()

	fmt.Printf("source is the following: %s\n", input.source)

	got := input.tokens

	want := []string{"https://github.com/rollenreiter/delorean", "junktext", "http://asahina.moe", "github.com/stompman34"}
	fmt.Printf("want following tokens:\n%s\n", want)

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	for i := range got {
		if got[i].content != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestTokenizeFile(t *testing.T) {
	Flags.SetFlags(FromTestfile)
	input := NewInput()
	input.Tokenize()

	fmt.Printf("source is the following: %s\n", input.source)
	got := input.tokens

	want := []string{
		"foo",
		"bar",
		"https://github.com/rollenreiter/delorean",
		"junktext",
		"http://asahina.moe",
		"github.com/stompman34",
	}

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	for i := range got {
		if got[i].content != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestGetUrls(t *testing.T) {
	Flags.SetFlags(FromTestURL)
	var wg sync.WaitGroup
	input := NewInput()
	input.Tokenize()
	input.GetUrls(&wg)
	got := input.validUrls
	fmt.Println(input)
	want := []string{
		"https://github.com/rollenreiter/delorean",
		"http://asahina.moe",
		"http://github.com/stompman34",
	}

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	Sort(got, false)
	for i := range got {
		if got[i].content != want[i] {
			t.Errorf("%q is not within %q", want, got)
		}
	}
}

func TestGetUrlsFile(t *testing.T) {
	Flags.SetFlags(FromTestURL)
	var wg sync.WaitGroup
	input := NewInput()
	input.Tokenize()
	fmt.Println(input.tokens)
	input.GetUrls(&wg)
	got := input.validUrls
	fmt.Println(got)
	want := []string{
		"https://github.com/rollenreiter/delorean",
		"http://asahina.moe",
		"http://github.com/stompman34",
	}

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	Sort(got, false)
	for i := range got {
		if got[i].content != want[i] {
			t.Errorf("%q is not within %q", want, got)
		}
	}
}
