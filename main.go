package main

import (
	"fmt"
)

func main() {
	flags := GetFlags()
	urls := GetUrl(*flags)

	final := linksToArchive(*flags, urls)

	if flags.silentFlag {
		for i := range final {
			fmt.Printf("%s\n", final[i])
		}
	} else {
		fmt.Printf("These are the links to the archives:\n")
		for i := range final {
			fmt.Printf("%s\n", final[i])
		}
	}
}

func linksToArchive(f cmdflags, sites []string) []string {
	archives := make([]string, len(sites))
	for i := range sites {
		site := sites[i]

		if !f.silentFlag {
			fmt.Printf("Archiving %s...\n", site)
		}

		url := fmt.Sprintf("https://web.archive.org/save/%s\n", site)
		archives[i] = ParseSingleLink(url)
	}
	return archives
}
