package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// GetTokens determines the URL source by reading f, then sets u.source accordingly and appends
// every token inside of the source to u.tokens.
func (u *urls) GetTokens(f *cmdflags) {
	// behaviour when reading from a file
	if f.fileFlag != "" {
		u.source = f.fileFlag
		file, err := os.Open(u.source)
		if err != nil {
			fmt.Printf("Failed reading from %s\n", u.source)
		}
		defer file.Close()
		s := bufio.NewScanner(file)
		s.Split(bufio.ScanWords)

		for s.Scan() {
			u.tokens = append(u.tokens, s.Text())
		}
	}

	// behaviour when reading from urlFlag
	if f.urlFlag != "" {
		u.source = "urlFlag"
		u.tokens = strings.Fields(f.urlFlag)
	}

	// behaviour when reading from stdin
	if f.urlFlag == "" && f.fileFlag == "" {
		u.source = "stdin"
		fmt.Println("Enter a URL to archive:")
		u.tokens = append(u.tokens, "")
		_, err := fmt.Scanln(&u.tokens[0])
		if err != nil {
			log.Fatal(err)
		}
	}
}

// GetUrls prepends "http://" to every non-URL string in u.tokens and sends a http GET
// request to each to determine if it is valid.
// Once validity is verified, the URL is added to u.validUrls. Invalid URLs are discarded.
func (u *urls) GetUrls(f *cmdflags) {
	preprocess := func(s *string) {
		if !strings.Contains(*s, "https://") || !strings.Contains(*s, "http://") {
			*s = fmt.Sprintf("http://%s", *s)
		}
	}

	switch u.source {
	case f.fileFlag:
		{
			fmt.Println("Validating URLs...")
			for _, url := range u.tokens {
				preprocess(&url)
				_, err := http.Get(url)
				if err != nil {
					return
				} else {
					u.validUrls = append(u.validUrls, url)
				}
			}
		}

	case "urlFlag":
		{
			fmt.Println("Validating URLs...")
			for _, url := range u.tokens {
				preprocess(&url)

				_, err := http.Get(url)
				if err != nil {
					fmt.Printf("Could not resolve \"%s\", skipping\n", url)
					u.results = append(u.results, fmt.Sprintf("UNARCHIVED: %s", url))
					return
				} else {
					u.validUrls = append(u.validUrls, url)
				}
			}
		}

	case "stdin":
		{
			fmt.Println("Validating URL...")
			preprocess(&u.tokens[0])
			_, err := http.Get((u.tokens[0]))
			if err != nil {
				log.Fatal(err)
			} else {
				u.validUrls[0] = u.tokens[0]
			}
		}
	}
}
