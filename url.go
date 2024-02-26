package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

// Tokenize determines the URL source by reading f, then sets u.source accordingly and appends
// every token inside of the source to u.tokens.
func (u *urls) Tokenize() {
	// behaviour when reading from a file
	switch {
	case Flags.fileFlag != "":
		{
			u.source = Flags.fileFlag
			file, err := os.Open(u.source)
			if err != nil {
				fmt.Printf("Failed reading from %s\n", u.source)
				file.Close()
				os.Exit(1)
			}
			defer file.Close()
			s := bufio.NewScanner(file)
			s.Split(bufio.ScanWords)

			for i := 0; s.Scan(); i++ {
				newToken := token{
					order:   i,
					content: s.Text(),
				}
				u.tokens = append(u.tokens, newToken)
			}
		}
	// behaviour when reading from urlFlag
	case Flags.urlFlag != "":
		{
			u.source = "urlFlag"
			urls := strings.Fields(Flags.urlFlag)
			for o, c := range urls {
				newToken := token{
					order:   o,
					content: c,
				}
				u.tokens = append(u.tokens, newToken)
			}
		}

	// behaviour when reading from stdin
	case Flags.urlFlag == "" && Flags.fileFlag == "":
		{
			u.source = "stdin"
			InterfaceInit(u)
		}
	}
}

func (u *urls) TokenizeSilent() {
	// behaviour when reading from a file
	switch {
	case Flags.fileFlag != "":
		{
			u.source = Flags.fileFlag
			file, err := os.Open(u.source)
			if err != nil {
				file.Close()
				os.Exit(1)
			}
			defer file.Close()
			s := bufio.NewScanner(file)
			s.Split(bufio.ScanWords)

			for i := 1; s.Scan(); i++ {
				newToken := token{
					order:   i,
					content: s.Text(),
				}
				u.tokens = append(u.tokens, newToken)
			}
		}

	// behaviour when reading from urlFlag
	case Flags.urlFlag != "":
		{
			u.source = "urlFlag"
			urls := strings.Fields(Flags.urlFlag)
			for o, c := range urls {
				newToken := token{
					order:   o,
					content: c,
				}
				u.tokens = append(u.tokens, newToken)
			}
		}

	// behaviour when reading from stdin
	case Flags.urlFlag == "" && Flags.fileFlag == "":
		{
			u.source = "stdin"
			InterfaceInit(u)
		}
	}
}

// preprocess every token in s so that they satisfy http.Get
func preprocess(s []token) []token {
	// make new array of tokens
	p := make([]token, len(s))

	for i, url := range s {
		// string less than 8 bytes cant possibly be a string that satisfies http.Get
		if len(url.content) <= 8 {
			p[i] = token{
				order:   url.order,
				content: fmt.Sprintf("http://%s", strings.ToLower(url.content)),
			}
			continue
		}
		if url.content[:7] == "http://" || url.content[:8] == "https://" {
			p[i] = token{
				order:   url.order,
				content: strings.ToLower(url.content),
			}
			continue
		}
		p[i] = token{
			order:   url.order,
			content: fmt.Sprintf("http://%s", strings.ToLower(url.content)),
		}
	}
	return p
}

// GetUrls prepends "http://" to every non-URL string in u.tokens and sends a http GET
// request to each to determine if it is valid.
// Once validity is verified, the URL is added to u.validUrls. Invalid URLs are discarded.
func (u *urls) GetUrls(wg *sync.WaitGroup) {
	switch u.source {
	case Flags.fileFlag:
		{
			fmt.Println("Validating URLs...")
			processedUrls := preprocess(u.tokens)
			for _, url := range processedUrls {
				wg.Add(1)
				go func(url token) {
					_, err := http.Get(url.content)
					if err != nil {
						wg.Done()
						return
					} else {
						u.validUrls = append(u.validUrls, url)
						wg.Done()
					}
				}(url)
			}
		}

	case "urlFlag":
		{
			fmt.Println("Validating URLs...")
			processedUrls := preprocess(u.tokens)
			for _, url := range processedUrls {
				wg.Add(1)
				go func(url token) {
					_, err := http.Get(url.content)
					if err != nil {
						fmt.Printf("Could not resolve \"%s\", skipping...\n", url.content)
						newToken := token{
							order:   url.order,
							content: fmt.Sprintf("UNARCHIVED: %s", url.content),
						}
						u.results = append(u.results, newToken)
						wg.Done()
						return
					} else {
						u.validUrls = append(u.validUrls, url)
						wg.Done()
					}
				}(url)
			}
		}
	}
	wg.Wait()
}

func (u *urls) GetUrlsSilent(wg *sync.WaitGroup) {
	switch u.source {
	case Flags.fileFlag:
		{
			processedUrls := preprocess(u.tokens)
			for _, url := range processedUrls {
				wg.Add(1)
				go func(url token) {
					_, err := http.Get(url.content)
					if err != nil {
						wg.Done()
						return
					} else {
						u.validUrls = append(u.validUrls, url)
						wg.Done()
					}
				}(url)
			}
		}

	case "urlFlag":
		{
			processedUrls := preprocess(u.tokens)
			for _, url := range processedUrls {
				wg.Add(1)
				go func(url token) {
					_, err := http.Get(url.content)
					if err != nil {
						newToken := token{
							order:   url.order,
							content: fmt.Sprintf("UNARCHIVED: %s", url.content),
						}
						u.results = append(u.results, newToken)
						wg.Done()
						return
					} else {
						u.validUrls = append(u.validUrls, url)
						wg.Done()
					}
				}(url)
			}
		}
	}
	wg.Wait()
}
