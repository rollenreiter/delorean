package main

import (
	"fmt"
)

func main() {
	flags := GetFlags()
	urls := GetUrl(*flags)

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
