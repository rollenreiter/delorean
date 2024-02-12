package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func (in *urls) Archive(f *cmdflags) {
	for i := range in.urls {
		fmt.Printf("Archiving %s...\n", in.urls[i])
		resp, err := http.Get(fmt.Sprintf("https://web.archive.org/save/%s", in.urls[i]))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		archive := resp.Request.URL.String()

		validOutput := regexp.MustCompile(`http.:\/\/web\.archive\.org\/web\/[0-9]{14}\/`)
		if !validOutput.Match([]byte(archive)) {
			fmt.Printf("\"%s\" can't be archived. Please try archiving it in your browser",
				in.urls[i])
			in.results = append(in.results, fmt.Sprintf("UNARCHIVED: %s", in.urls[i]))
		} else {
			in.results = append(in.results, archive)
		}
	}
}
