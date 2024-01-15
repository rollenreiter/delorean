package main

import (
	"fmt"
	"regexp"
)

func main() {
	flags := GetFlags()
	urls := GetUrl(*flags)

	// final := linksToArchive(*flags, urls)
	linksToArchiveMutate(*flags, urls)

	if flags.silentFlag {
		for i := range urls {
			fmt.Printf("%s\n", urls[i])
		}
	} else {
		fmt.Printf("These are the links to the archives:\n")
		for i := range urls {
			fmt.Printf("%s\n", urls[i])
		}
	}
}

// deprecated and (potentially) slower
func linksToArchive(f cmdflags, sites []string) []string {
	archives := make([]string, len(sites))
	for i := range sites {
		var err error
		site := sites[i]

		if !f.silentFlag {
			fmt.Printf("Archiving %s...\n", site)
		}

		url := fmt.Sprintf("https://web.archive.org/save/%s\n", site)
		archives[i], err = ParseSingleLink(url)

		if err != nil {
			fmt.Println(err)
		}
	}
	return archives
}

func linksToArchiveMutate(f cmdflags, sites []string) {
	validInput := regexp.MustCompile(`[-a-zA-Z0-9@:%_\+.~#?&//=]{2,256}\.[a-z]{2,13}\b(\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?`)
	for i := range sites {
		if !validInput.Match([]byte(sites[i])) {
			fmt.Printf("\"%s\" is not a valid URL.\n", sites[i])
			sites[i] = ""
			break
		}
		site := fmt.Sprintf("https://web.archive.org/save/%s\n", sites[i])

		if !f.silentFlag {
			fmt.Printf("Archiving %s...\n", sites[i])
		}
		ParseSingleLinkMutate(&site)
		sites[i] = site
	}
}
