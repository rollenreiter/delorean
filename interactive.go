package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	Colors  = InitColors()
	version = "v0.1.0"
	Greeter = fmt.Sprintf("Welcome to DeLorean %s\nType h for help\n", version)
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
		case "d":
			InterfaceDelete(u)
		case "h":
			InterfaceHelp()
		case "l":
			InterfaceList(u)
		case "w":
			InterfaceArchive(u)
		default:
			fmt.Printf("Unknown command.\n\n")
		}
	}
}

func InterfaceHelp() {
	fmt.Printf("Keybinds:\n")
	fmt.Printf("a - Add a URL to the archive list\n")
	fmt.Printf("l - List all links in the archive list\n")
	fmt.Printf("d - Remove a URL from the archive list\n")
	fmt.Printf("w - Archive all URLs in archive list\n")
	fmt.Printf("q - Quit interactive interface\n")
	fmt.Println()
}

func InterfaceAdd(u *urls) {
	fmt.Println("Enter the URLs to add to the archive list:")
	fmt.Print("[" + fmt.Sprint(len(u.tokens)) + "] + ")
	var new string
	fmt.Scanln(&new)
	fmt.Println(new)
	if new == "\n" {
		fmt.Printf("Aborted.\n\n")
	} else {
		add := strings.Split(new, " ")
		u.tokens = append(u.tokens, add...)
		fmt.Printf("Successfully added to archive list\n\n")
	}
}

func InterfaceDelete(u *urls) {
	fmt.Println("Which URL do you want to remove?:")
	for i, s := range u.tokens {
		fmt.Printf("(%d) %s\n", i+1, s)
	}
	fmt.Print("[" + fmt.Sprint(len(u.tokens)) + "] - ")
	var selection string
	fmt.Scanln(&selection)
	deleted, err := strconv.Atoi(selection)
	switch {
	case err != nil:
		fmt.Printf("Please enter a number.\n\n")
	case deleted > len(u.tokens):
		fmt.Printf("This URL is not in the list.\n\n")
	default:
		{
			fmt.Printf("Successfully deleted \"%s\" from archive list.\n\n", u.tokens[deleted-1])
			new := make([]string, 0, len(u.tokens)-1)
			for i, s := range u.tokens {
				if i == deleted-1 {
					continue
				}
				new = append(new, s)
			}
			u.tokens = new
		}
	}
}

func InterfaceList(u *urls) {
	if len(u.tokens) == 0 {
		fmt.Printf("The archive list is currently empty.\n\n")
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
	if len(u.tokens) != 0 {
		u.ArchiveInter(&wg)
		u.Finish()
	}
	fmt.Println("Nothing to archive.")
	fmt.Println()
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
				fmt.Printf("Could not resolve \"%s\", skipping...\n", url)
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
