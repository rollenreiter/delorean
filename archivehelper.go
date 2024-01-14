package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

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
