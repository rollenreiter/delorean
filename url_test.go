package main

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	flags := cmdflags{
		urlFlag: "github.com/rollenreiter/delorean junktext archive.org",
	}
	input := NewInput()
	input.GetTokens(&flags)
	got := input.tokens
	want := []string{"github.com/rollenreiter/delorean", "junktext", "archive.org"}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestTokenizeFile(t *testing.T) {
	flags := cmdflags{
		urlFlag:  "",
		fileFlag: "testfile",
	}
	input := NewInput()
	input.GetTokens(&flags)
	got := input.tokens
	want := []string{
		"foo",
		"bar",
		"github.com/rollenreiter/delorean",
		"junktext",
		"archive.org",
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestGetUrls(t *testing.T) {
	flags := cmdflags{
		urlFlag:  "github.com/rollenreiter/delorean junktext archive.org",
		fileFlag: "",
	}
	input := NewInput()
	input.GetTokens(&flags)
	input.GetUrls(&flags)
	got := input.validUrls
	want := []string{"http://github.com/rollenreiter/delorean", "http://archive.org"}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestGetUrlsFile(t *testing.T) {
	flags := cmdflags{
		urlFlag:  "",
		fileFlag: "testfile",
	}
	input := NewInput()
	input.GetTokens(&flags)
	input.GetUrls(&flags)
	got := input.validUrls
	want := []string{
		"http://github.com/rollenreiter/delorean",
		"http://archive.org",
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
