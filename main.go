package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/rollenreiter/delorean/parser"
)

type cmdflags struct {
	silentFlag bool
	urlFlag    string
}

func main() {
	flags := getFlags()
	urls := getUrl(*flags)

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
		archives[i] = parser.ParseSingleLink(url)
	}
	return archives
}

// getting urls and saving them as a string slice of urls
func getUrl(f cmdflags) []string {
	u := []string{""}
	if f.urlFlag == "" {
		fmt.Println("Enter a link to archive:")
		_, err := fmt.Scanln(&u[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		urls := f.urlFlag
		u = strings.Fields(urls)
	}
	return u
}

// responsible for parsing command line flags
func getFlags() *cmdflags {
	var f cmdflags
	flag.BoolVar(&f.silentFlag, "s", false, "Supress all output except for the final link to the archive; useful for scripting")
	flag.StringVar(&f.urlFlag, "u", "", "Declare URLs to archive as a single string or a space-seperated sequence of strings; useful for non-interactive use, scripting and archiving multiple URLs at once")
	flag.Parse()
	return &f
}
