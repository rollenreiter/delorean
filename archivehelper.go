package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// deprecated and (probably) slower
func ParseSingleLink(url string) (string, error) {
	var err error
	trimmed := strings.Trim(url, "\n")
	validInput := regexp.MustCompile(`[-a-zA-Z0-9@:%_\+.~#?&//=]{2,256}\.[a-z]{2,13}\b(\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?`)
	if !validInput.Match([]byte(trimmed)) {
		err = fmt.Errorf("error: %s is not a real URL", url)
		log.Fatal(err)
	}
	resp, err := http.Get(trimmed)
	if err != nil {
		log.Fatal(err)
	}
	finalurl := resp.Request.URL.String()
	validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
	if !validOutput.Match([]byte(finalurl)) {
		err = fmt.Errorf("error: couldn't get archive URL. please try archiving %s in your browser", url)
	}
	return finalurl, err
}

// directly mutates the url passed into it, changed to archive url
func ParseSingleLinkMutate(url *string) {
	*url = strings.Trim(*url, "\n")
	validInput := regexp.MustCompile(`[-a-zA-Z0-9@:%_\+.~#?&//=]{2,256}\.[a-z]{2,13}\b(\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?`)
	if !validInput.Match([]byte(*url)) {
		fmt.Printf("\"%s\" is not a valid URL.\n", *url)
	}
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	*url = resp.Request.URL.String()
	validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
	if !validOutput.Match([]byte(*url)) {
		fmt.Printf("the archive URL is not valid. please try archiving %s in your browser", *url)
	}
}
