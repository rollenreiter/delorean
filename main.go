package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rollenreiter/delorean/parser"
)

type cmdflags struct {
	silentFlag bool
	urlFlag    string
}

func main() {
	flags := getFlags()
	url, err := getUrl(*flags)
	if err != nil {
		log.Fatal(err)
	}
	final := linkToArchive(*flags, url)

	if flags.silentFlag {
		fmt.Println(final)
	} else {
		fmt.Printf("This is the link to the archive:\n%s\n", final)
	}
}

func linkToArchive(f cmdflags, site string) string {
	if !f.silentFlag {
		fmt.Printf("Archiving %s...\n", site)
	}
	url := fmt.Sprintf("https://web.archive.org/save/%s\n", site)
	archive := parser.ParseSingleLink(url)
	return archive
}

func getUrl(f cmdflags) (string, error) {
	var err error
	var url string
	if f.urlFlag == "" {
		fmt.Println("Give me a link to archive:")
		_, err := fmt.Scanln(&url)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		url = f.urlFlag
	}
	return url, err
}

func getFlags() *cmdflags {
	var f cmdflags
	flag.BoolVar(&f.silentFlag, "s", false, "Supress all output except for the final link to the archive; useful for scripting")
	flag.BoolVar(&f.silentFlag, "silent", false, "Supress all output except for the final link to the archive; useful for scripting")
	flag.StringVar(&f.urlFlag, "u", "", "Declare a URL to archive; useful for scripting and archiving multiple URLs at once")
	flag.Parse()
	return &f
}
