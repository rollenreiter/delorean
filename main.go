package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rollenreiter/delorean/parser"
)

func main() {
	url, err := parseUrlFromCmdline()
	if err != nil {
		log.Fatal(err)
	}
	final := getLinkToArchive(url)
	fmt.Printf("This is the link to the archive:\n%s\n", final)
}

func getLinkToArchive(site string) string {
	fmt.Printf("Archiving %s...\n", site)
	url := fmt.Sprintf("https://web.archive.org/save/%s\n", site)
	archive := parser.ParseSingleLink(url)
	return archive
}

func parseUrlFromCmdline() (string, error) {
	var input []string
	var err error
	input = os.Args
	flag.Parse()
	// throw error if no url is provided
	if input[len(input)-1] == input[0] || input[1] == "" {
		err = fmt.Errorf("please supply a url to archive")
	}
	url := input[len(input)-1]
	return url, err
}
