package main

import (
	"testing"
)

func Copy(b *testing.B) {
	sites := []string{
		"milk.com",
	}
	flags := cmdflags{
		silentFlag: false,
		urlFlag:    "",
	}

	for n := 0; n < b.N; n++ {
		linksToArchive(flags, sites)
	}
}

func Mutate(b *testing.B) {
	sites := []string{
		"milk.com",
	}
	flags := cmdflags{
		silentFlag: true,
	}

	for n := 0; n < b.N; n++ {
		linksToArchiveMutate(flags, sites)
	}
}
