package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	Colors  = InitColors()
	version = "v0.1.0"
	Greeter = fmt.Sprintf("Welcome to DeLorean %s\n", version)
)

func InterfaceInit(u *urls) {
	fmt.Print(Greeter)

	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("[" + fmt.Sprint(len(u.tokens)) + "] > ")
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		switch s {
		case "q":
			os.Exit(1)
		case "a":
			InterfaceAdd(u)
		case "h":
			InterfaceHelp()
		case "l":
			InterfaceList(u)
		case "w":
			InterfaceArchive(u)
		}
	}
}

func InterfaceHelp() {
	fmt.Printf("Keybinds:\n")
	fmt.Printf("q\tQuit interactive interface\n")
	fmt.Printf("l\tList all links in the archive list\n")
	fmt.Printf("a\tAdd a URL to the archive list\n")
	fmt.Printf("w\tArchive all URLs in archive list\n")
	fmt.Println()
}

func InterfaceAdd(u *urls) {
	fmt.Println("Enter a URL to add to the archive list:")
	fmt.Print("[" + fmt.Sprint(len(u.tokens)) + "] + ")
	var new string
	fmt.Scanln(&new)
	u.tokens = append(u.tokens, new)
	fmt.Printf("Successfully added %s to archive list\n", new)
	fmt.Println()
}

func InterfaceList(u *urls) {
	if len(u.tokens) == 0 {
		fmt.Printf("The archive list is currently empty\n\n")
	} else {
		fmt.Printf("These links will be archived on write:\n")
		for i, s := range u.tokens {
			fmt.Printf("(%d) %s\n", i+1, s)
		}
		fmt.Println()
	}
}

func InterfaceArchive(u *urls) {
	var wg sync.WaitGroup
	u.ArchiveInter(&wg)
	u.Finish()
}

func (u *urls) ArchiveInter(wg *sync.WaitGroup) {
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
	wg.Wait()

	fmt.Printf("\nArchiving all URLs. Depending on the Internet Archive's traffic, this may take a long time.\n")
	for _, url := range u.validUrls {
		wg.Add(1)
		go func(url string) {
			fmt.Printf("Archiving %s...\n", url)
			resp, err := http.Get(fmt.Sprintf("https://web.archive.org/save/%s", url))
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			archive := resp.Request.URL.String()

			validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
			if !validOutput.Match([]byte(archive)) {
				fmt.Printf("\"%s\" couldn't be archived. This may be due to the website being blacklisted from archive.org. For more information, please try archiving it in your browser\n",
					url)
				u.results = append(u.results, fmt.Sprintf("UNARCHIVED: %s", url))
				wg.Done()
			} else {
				u.results = append(u.results, archive)
				wg.Done()
			}
		}(url)
	}
	wg.Wait()
}

func (u urls) Finish() {
	fmt.Println("\nSUCCESS! These are the links to the archives:")
	for _, s := range u.results {
		fmt.Println(s)
	}
	os.Exit(0)
}
