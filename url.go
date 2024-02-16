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
func (u *urls) Tokenize(f *cmdflags) {
	// behaviour when reading from a file
	if f.fileFlag != "" {
		u.source = f.fileFlag
		file, err := os.Open(u.source)
		if err != nil {
			fmt.Printf("Failed reading from %s\n", u.source)
			file.Close()
			os.Exit(1)
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
		InterfaceInit(u)
	}
}

func (u *urls) TokenizeSilent(f *cmdflags) {
	// behaviour when reading from a file
	if f.fileFlag != "" {
		u.source = f.fileFlag
		file, err := os.Open(u.source)
		if err != nil {
			fmt.Printf("Failed reading from %s\n", u.source)
			file.Close()
			os.Exit(1)
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
	// TODO: make actual interactive mode
	if f.urlFlag == "" && f.fileFlag == "" {
		u.source = "stdin"
		InterfaceInit(u)
	}
}

// GetUrls prepends "http://" to every non-URL string in u.tokens and sends a http GET
// request to each to determine if it is valid.
// Once validity is verified, the URL is added to u.validUrls. Invalid URLs are discarded.
func (u *urls) GetUrls(f *cmdflags, wg *sync.WaitGroup) {
	preprocess := func(s []string) []string {
		p := make([]string, len(s))
		for i, url := range s {
			if len(url) <= 8 {
				p[i] = fmt.Sprintf("http://%s", url)
				continue
			}
			if url[:7] == "http://" || url[:8] == "https://" {
				p[i] = url
				continue
			}
			p[i] = fmt.Sprintf("http://%s", url)
		}
		return p
	}

	switch u.source {
	case f.fileFlag:
		{
			fmt.Println("Validating URLs...")
			processedUrls := preprocess(u.tokens)
			for _, url := range processedUrls {
				wg.Add(1)
				go func(url string) {
					_, err := http.Get(url)
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
				go func(url string) {
					_, err := http.Get(url)
					if err != nil {
						fmt.Printf("Could not resolve \"%s\", skipping\n", url)
						u.results = append(u.results, fmt.Sprintf("UNARCHIVED: %s", url))
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

func (u *urls) GetUrlsSilent(f *cmdflags, wg *sync.WaitGroup) {
	preprocess := func(s []string) []string {
		p := make([]string, len(s))
		for i, url := range s {
			if len(url) <= 8 {
				p[i] = fmt.Sprintf("http://%s", url)
				continue
			}
			if url[:7] == "http://" || url[:8] == "https://" {
				p[i] = url
				continue
			}
			p[i] = fmt.Sprintf("http://%s", url)
		}
		return p
	}

	switch u.source {
	case f.fileFlag:
		{
			processedUrls := preprocess(u.tokens)
			for _, url := range processedUrls {
				wg.Add(1)
				go func(url string) {
					_, err := http.Get(url)
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
				go func(url string) {
					_, err := http.Get(url)
					if err != nil {
						u.results = append(u.results, fmt.Sprintf("UNARCHIVED: %s", url))
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
