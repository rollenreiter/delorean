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
		silentFlag: true,
		alphaFlag:  false,
	}
	return f
}

func FromTestfile() cmdflags {
	f := cmdflags{
		urlFlag:    "",
		fileFlag:   "testfile",
		silentFlag: true,
		alphaFlag:  false,
	}
	return f
}

func TestTokenizeUrl(t *testing.T) {
	fmt.Println("Running TestTokenizeUrl...")
	Flags.SetFlags(FromTestURL)
	input := NewInput()
	input.Tokenize()

	got := input.tokens

	want := []string{"https://github.com/rollenreiter/delorean", "junktext", "http://asahina.moe", "github.com/stompman34"}

	if len(got) == 0 {
		t.Errorf("slice %q is empty, wanted %q", got, want)
	}
	for i := range got {
		if got[i].content != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
	fmt.Println("Done.")
}

func TestTokenizeFile(t *testing.T) {
	fmt.Println("Running TestTokenizeFile")
	Flags.SetFlags(FromTestfile)
	input := NewInput()
	input.Tokenize()

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
	fmt.Println("Done.")
}

func TestGetUrls(t *testing.T) {
	fmt.Println("Running TestGetUrls...")
	Flags.SetFlags(FromTestURL)
	var wg sync.WaitGroup
	input := NewInput()
	input.Tokenize()
	input.GetUrls(&wg)
	got := input.validUrls
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
	fmt.Println("Done.")
}

func TestGetUrlsFile(t *testing.T) {
	fmt.Println("Running TestGetUrlsFile...")
	Flags.SetFlags(FromTestURL)
	var wg sync.WaitGroup
	input := NewInput()
	input.Tokenize()
	input.GetUrls(&wg)
	got := input.validUrls
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
	fmt.Println("Done.")
}
