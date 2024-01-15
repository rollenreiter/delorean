package main

import (
	"fmt"
)

func main() {
	flags := GetFlags()
	var urls []string
	if flags.fileFlag == "" {
		urls = GetUrl(*flags)
	} else {
		urls = GetUrlFile(*flags)
	}

	getLinkToArchive(*flags, urls)

	if flags.silentFlag {
		for i := range urls {
			fmt.Printf("%s\n", urls[i])
		}
	} else {
		fmt.Printf("These are the links to the archives:\n")
		for i := range urls {
			fmt.Printf("%s\n", urls[i])
		}
	}
}
