package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func requestSingleLink(url *string) {
	*url = strings.Trim(*url, "\n")
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	archive := resp.Request.URL.String()
	validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
	if !validOutput.Match([]byte(archive)) {
		*url = fmt.Sprintf("The archive URL is not valid. Please try this link:\n%s\n in your browser", *url)
	} else {
		*url = archive
	}
}

func getLinkToArchive(f cmdflags, sites []string) {
	preprocess := func(s *string) {
		if !strings.Contains(*s, "https://") || !strings.Contains(*s, "http://") {
			*s = fmt.Sprintf("http://%s", *s)
		}
	}
	for i := range sites {
		preprocess(&sites[i])
		fmt.Printf("Checking validity of %s...\n", sites[i])
		_, err := http.Get((sites[i]))
		if err != nil {
			if !f.silentFlag {
				fmt.Printf("\"%s\" is not a valid URL, skipping.\n", sites[i])
			}
			sites[i] = fmt.Sprintf("UNARCHIVED: %s", sites[i])
		} else {
			site := fmt.Sprintf("https://web.archive.org/save/%s\n", sites[i])

			if !f.silentFlag {
				fmt.Printf("Archiving %s...\n", sites[i])
			}
			requestSingleLink(&site)
			sites[i] = site
		}
	}
}

func getLinkToArchiveFile(f cmdflags, sites []string) {
	validInput := regexp.MustCompile(`[-a-zA-Z0-9@:%_\+.~#?&//=]{2,256}\.[a-z]{2,13}\b(\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?`)
	for i := range sites {
		if validInput.Match([]byte(sites[i])) {
			site := fmt.Sprintf("https://web.archive.org/save/%s\n", sites[i])

			if !f.silentFlag {
				fmt.Printf("Archiving %s...\n", sites[i])
			}
			requestSingleLink(&site)
			sites[i] = site
		}
	}
}
