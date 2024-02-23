package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
)

// Archive archives every URL in u.urls by sending a http GET request to the Internet
// Archive and waiting for the final redirect, being the archived web page.
// Should the final redirect's URL not match the usual pattern of an archived web page,
// Archive will assume an issue on the Internet Archive's end and prompt the user to
// archive the URL that caused the issue in their browser.
func (u *urls) Archive(f *cmdflags, wg *sync.WaitGroup) {
	fmt.Printf("\nArchiving all URLs. Depending on the Internet Archive's traffic, this may take a long time.\n")
	for _, url := range u.validUrls {
		wg.Add(1)
		go func(url token) {
			fmt.Printf("Archiving %s\n", url.content)
			resp, err := http.Get(fmt.Sprintf("https://web.archive.org/save/%s", url.content))
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			archive := resp.Request.URL.String()

			validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
			if !validOutput.Match([]byte(archive)) {
				fmt.Printf("\"%s\" couldn't be archived. This may be due to the website being blacklisted from archive.org. For more information, please try archiving it in your browser\n",
					url.content)
				newToken := token{
					order:   url.order,
					content: fmt.Sprintf("UNARCHIVED: %s", url.content),
				}
				u.results = append(u.results, newToken)
				wg.Done()
			} else {

				newToken := token{
					order:   url.order,
					content: archive,
				}
				u.results = append(u.results, newToken)
				wg.Done()
			}
		}(url)
	}
	wg.Wait()
}

func (u *urls) ArchiveSilent(f *cmdflags, wg *sync.WaitGroup) {
	for _, url := range u.validUrls {
		wg.Add(1)
		go func(url token) {
			resp, err := http.Get(fmt.Sprintf("https://web.archive.org/save/%s", url.content))
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			archive := resp.Request.URL.String()

			validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
			if !validOutput.Match([]byte(archive)) {
				newToken := token{
					order:   url.order,
					content: fmt.Sprintf("UNARCHIVED: %s", url.content),
				}
				u.results = append(u.results, newToken)
				wg.Done()
			} else {

				newToken := token{
					order:   url.order,
					content: archive,
				}
				u.results = append(u.results, newToken)
				wg.Done()
			}
		}(url)
	}
	wg.Wait()
}

func (u *urls) ArchiveInter(wg *sync.WaitGroup) {
	fmt.Println("Validating URLs...")
	processedUrls := preprocess(u.tokens)
	for _, url := range processedUrls {
		wg.Add(1)
		go func(url token) {
			_, err := http.Get(url.content)
			if err != nil {
				fmt.Printf("Could not resolve \"%s\", skipping...\n", url.content)
				unarchivedToken := token{
					order:   url.order,
					content: fmt.Sprintf("UNARCHIVED: %s", url.content),
				}
				u.results = append(u.results, unarchivedToken)
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
		go func(url token) {
			resp, err := http.Get(fmt.Sprintf("https://web.archive.org/save/%s", url.content))
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			archive := resp.Request.URL.String()

			validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
			if !validOutput.Match([]byte(archive)) {
				fmt.Printf("\"%s\" couldn't be archived. This may be due to the website being blacklisted from archive.org. For more information, please try archiving it in your browser\n",
					url.content)
				unarchivedToken := token{
					order:   url.order,
					content: fmt.Sprintf("UNARCHIVED: %s", url.content),
				}
				u.results = append(u.results, unarchivedToken)
				wg.Done()
			} else {
				archivedToken := token{
					order:   url.order,
					content: archive,
				}
				u.results = append(u.results, archivedToken)
				wg.Done()
			}
		}(url)
	}
	wg.Wait()
}
