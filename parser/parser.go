package parser

import (
	"log"
	"net/http"
)

func ParseSingleLink(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	finalurl := resp.Request.URL.String()
	return finalurl
}
