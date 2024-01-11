package parser

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func Parse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// resp, err := os.Open("test.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var links []string
	var link func(*html.Node)
	// anonymous function
	link = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			link(c)
		}
	}
	link(doc)
	for _, l := range links {
		fmt.Println(l)
	}
}
