package parser

import (
	"log"
	"net/http"
	"strings"
)

func ParseSingleLink(url string) string {
	resp, err := http.Get(strings.Trim(url, "\n"))
	if err != nil {
		log.Fatal(err)
	}
	finalurl := resp.Request.URL.String()
	return finalurl
}
