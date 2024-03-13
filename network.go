package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"
)

// Archive archives every URL in u.urls by sending a http GET request to the Internet
// Archive and waiting for the final redirect, being the archived web page.
// Should the final redirect's URL not match the usual pattern of an archived web page,
// Archive will assume an issue on the Internet Archive's end and prompt the user to
// archive the URL that caused the issue in their browser.
func (u *urls) Archive(wg *sync.WaitGroup) {
	if !Flags.silentFlag {
		fmt.Printf("\nArchiving all URLs. Depending on the Internet Archive's traffic, this may take a while.\n")
	}
	for _, parsedurl := range u.validUrls {
		wg.Add(1)
		go func(parsedurl token) {
			resp, err := http.Get(fmt.Sprintf("https://web.archive.org/save/%s", parsedurl.content))
			if err != nil {
				fmt.Printf("Couldn't connect to the Internet Archive while trying to archive %s.\n Please check your internet connection.\n", parsedurl.content)
				return
			}
			defer resp.Body.Close()
			archive := resp.Request.URL.String()

			validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
			if !validOutput.Match([]byte(archive)) {
				if !Flags.silentFlag {
					fmt.Printf("\"%s\" couldn't be archived. This may be due to the website being blacklisted from archive.org. For more information, please try archiving it in your browser\n",
						parsedurl.content)
				}
				newToken := token{
					order:   parsedurl.order,
					content: fmt.Sprintf("UNARCHIVED: %s", parsedurl.content),
				}
				u.results = append(u.results, newToken)
				wg.Done()
			} else {

				newToken := token{
					order:   parsedurl.order,
					content: archive,
				}
				u.results = append(u.results, newToken)
				wg.Done()
			}
		}(parsedurl)
	}
	wg.Wait()
	if !Flags.silentFlag {
		fmt.Println()
	}
}

// u.ArchiveInter() is the same as calling u.GetUrls and u.Archive, with some
// adjustments for the interactive CLI
func (u *urls) ArchiveInter(wg *sync.WaitGroup) {
	fmt.Println("Validating URLs...")
	processedUrls := Preprocess(u.tokens)
	for _, parsedtoken := range processedUrls {
		wg.Add(1)
		go func(parsedtoken token) {
			_, err := http.Get(parsedtoken.content)
			if err != nil {
				fmt.Printf("Could not resolve \"%s\", skipping...\n", parsedtoken.content)
				unarchivedToken := token{
					order:   parsedtoken.order,
					content: fmt.Sprintf("UNARCHIVED: %s", parsedtoken.content),
				}
				u.results = append(u.results, unarchivedToken)
				wg.Done()
				return
			} else {
				u.validUrls = append(u.validUrls, parsedtoken)
				wg.Done()
			}
		}(parsedtoken)
	}
	wg.Wait()

	fmt.Printf("\nArchiving all URLs. Depending on the Internet Archive's traffic, this may take a long time.\n")
	for _, parsedurl := range u.validUrls {
		wg.Add(1)
		go func(parsedurl token) {
			resp, err := http.Get(fmt.Sprintf("https://web.archive.org/save/%s", parsedurl.content))
			if err != nil {
				fmt.Printf("%sERROR:%s ", Error, Escape)
				fmt.Printf("Couldn't connect to archive.org while trying to archive %s.\n Please check your Internet connection.\n", parsedurl.content)
				return
			}
			defer resp.Body.Close()
			archive := resp.Request.URL.String()

			validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
			if !validOutput.Match([]byte(archive)) {
				fmt.Printf("%sWARNING:%s ", Warning, Escape)
				fmt.Printf("\"%s\" couldn't be archived. This may be due to the website being blacklisted from archive.org. For more information, please try archiving it in your browser\n",
					parsedurl.content)
				unarchivedToken := token{
					order:   parsedurl.order,
					content: fmt.Sprintf("UNARCHIVED: %s", parsedurl.content),
				}
				u.results = append(u.results, unarchivedToken)
				wg.Done()
			} else {
				archivedToken := token{
					order:   parsedurl.order,
					content: archive,
				}
				u.results = append(u.results, archivedToken)
				wg.Done()
			}
		}(parsedurl)
	}
	wg.Wait()
}

// CheckConnection will check whether or not a connection to the Internet Archive can be established.
// By default, the function will timeout after 8 seconds.
// TODO: make a timeout flag
func CheckConnection(wg *sync.WaitGroup) {
	if !Flags.silentFlag {
		fmt.Printf("Checking availability...\n")
	}
	wg.Add(1)
	netfailure := make(chan bool, 1)
	go func() {
		_, err := http.Get("http://web.archive.org")
		if err != nil {
			netfailure <- true
		} else {
			netfailure <- false
		}
	}()
	select {
	case failed := <-netfailure:
		if failed {
			fmt.Printf("%sERROR:%s ", Error, Escape)
			fmt.Println("Couldn't connect to the Internet Archive. Please check your internet connection.")
			os.Exit(1)
		} else {
			wg.Done()
		}
	case <-time.After(8 * time.Second):
		fmt.Printf("%sERROR:%s ", Error, Escape)
		fmt.Println("Timed out while trying to connect to the Internet Archive. Please check your internet connection.")
		os.Exit(1)
	}
	wg.Wait()
}
