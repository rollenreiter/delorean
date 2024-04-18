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
	case Flags.fromFile != "":
		{
			u.source = Flags.fromFile
			file, err := os.Open(u.source)
			if err != nil {
				fmt.Printf("%s|x|%s ", Error, Escape)
				fmt.Printf("Cannot %s\n", err)
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
	case Flags.fromString != "":
		{
			u.source = "urlFlag"
			flagtokens := strings.Fields(Flags.fromString)
			for o, c := range flagtokens {
				newToken := token{
					order:   o,
					content: c,
				}
				u.tokens = append(u.tokens, newToken)
			}
		}

	// behaviour when reading from stdin
	case Flags.fromString == "" && Flags.fromFile == "":
		{
			u.source = "stdin"
			InterfaceInit(u)
		}
	}
}

// GetUrls prepends "http://" to every non-URL string in u.tokens and sends a http GET
// request to each to determine if it is valid.
// Once validity is verified, the URL is added to u.validUrls. Invalid URLs are discarded.
func (u *urls) GetUrls(wg *sync.WaitGroup) {
	switch u.source {
	case Flags.fromFile:
		{
			processedUrls := Preprocess(u.tokens)
			for _, parsedtoken := range processedUrls {
				wg.Add(1)
				go func(parsedtoken token) {
					_, err := http.Get(parsedtoken.content)
					if err != nil {
						wg.Done()
						return
					} else {
						u.validUrls = append(u.validUrls, parsedtoken)
						wg.Done()
					}
				}(parsedtoken)
			}
		}

	case "urlFlag":
		{
			processedUrls := Preprocess(u.tokens)
			for _, parsedtoken := range processedUrls {
				wg.Add(1)
				go func(parsedtoken token) {
					_, err := http.Get(parsedtoken.content)
					if err != nil {
						if !Flags.silentFlag {
							fmt.Printf("%s|!|%s ", Warning, Escape)
							fmt.Printf("Could not resolve \"%s\", skipping...\n", parsedtoken.content)
						}
						newToken := token{
							order:   parsedtoken.order,
							content: fmt.Sprintf("UNARCHIVED: %s", parsedtoken.content),
						}
						u.results = append(u.results, newToken)
						wg.Done()
						return
					} else {
						u.validUrls = append(u.validUrls, parsedtoken)
						wg.Done()
					}
				}(parsedtoken)
			}
		}
	}
	wg.Wait()
}
