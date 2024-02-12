package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func (in *urls) GetTokens(f *cmdflags) {
	// behaviour when reading from a file
	if f.fileFlag != "" {
		in.source = f.fileFlag
		file, err := os.Open(in.source)
		if err != nil {
			fmt.Printf("Failed reading from %s\n", in.source)
		}
		defer file.Close()
		s := bufio.NewScanner(file)
		s.Split(bufio.ScanWords)

		for s.Scan() {
			in.tokens = append(in.tokens, s.Text())
		}
	}

	// behaviour when reading from urlFlag
	if f.urlFlag != "" {
		in.source = "urlFlag"
		in.tokens = strings.Fields(f.urlFlag)
	}

	// behaviour when reading from stdin
	if f.urlFlag == "" && f.fileFlag == "" {
		in.source = "stdin"
		fmt.Println("Enter a URL to archive:")
		in.tokens = append(in.tokens, "")
		_, err := fmt.Scanln(&in.tokens[0])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (in *urls) GetUrls(f *cmdflags) {
	preprocess := func(s *string) {
		if !strings.Contains(*s, "https://") || !strings.Contains(*s, "http://") {
			*s = fmt.Sprintf("http://%s", *s)
		}
	}
	switch in.source {

	case f.fileFlag:
		{
			fmt.Println("Validating URLs...")
			for i := range in.tokens {
				preprocess(&in.tokens[i])

				_, err := http.Get((in.tokens[i]))
				if err != nil {
					i++
				} else {
					in.urls = append(in.urls, in.tokens[i])
				}
			}
		}

	case "urlFlag":
		{
			fmt.Println("Validating URLs...")
			for i := range in.tokens {
				preprocess(&in.tokens[i])

				_, err := http.Get((in.tokens[i]))
				if err != nil {
					fmt.Printf("Could not resolve \"%s\", skipping\n", in.tokens[i])
					in.results = append(in.results, fmt.Sprintf("UNARCHIVED: %s", in.tokens[i]))
					i++
				} else {
					in.urls = append(in.urls, in.tokens[i])
				}
			}
		}

	case "stdin":
		{
			fmt.Println("Validating URL...")
			preprocess(&in.tokens[0])
			_, err := http.Get((in.tokens[0]))
			if err != nil {
				log.Fatal(err)
			} else {
				in.urls[0] = in.tokens[0]
			}
		}
	}
}
