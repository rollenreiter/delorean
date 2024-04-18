package main

import (
	"fmt"
	"os"
	"strings"
)

// Basic Bubble Sort for token types
func Sort(u []token, alphabetical bool) {
	swap := func(u []token, p1 int, p2 int) {
		temp := u[p2]
		u[p2] = u[p1]
		u[p1] = temp
	}
	swapped := true
	for swapped {
		swapped = false
		for i := range u {
			switch {
			case i == 0:
				continue
			case !alphabetical:
				{
					if u[i-1].order > u[i].order {
						swap(u, i-1, i)
						swapped = true
					}
				}
			default:
				{
					if u[i-1].content[0] > u[i].content[0] {
						swap(u, i-1, i)
						swapped = true

					}
				}
			}
		}
	}
}

// Preprocess every token in s so that they satisfy http.Get
func Preprocess(s []token) []token {
	// make new array of tokens
	p := make([]token, len(s))

	for i, url := range s {
		// string less than 8 bytes cant possibly be a string that satisfies http.Get
		if len(url.content) <= 8 {
			p[i] = token{
				order:   url.order,
				content: fmt.Sprintf("http://%s", strings.ToLower(url.content)),
			}
			continue
		}
		if url.content[:7] == "http://" || url.content[:8] == "https://" {
			p[i] = token{
				order:   url.order,
				content: strings.ToLower(url.content),
			}
			continue
		}
		p[i] = token{
			order:   url.order,
			content: fmt.Sprintf("http://%s", strings.ToLower(url.content)),
		}
	}
	return p
}

func IsMultifile() bool {
	var files int
	for i, s := range os.Args {
		switch {
		case i < 2:
			continue
		case files >= 1:
			return true
		case s[0] == '-':
			continue
		default:
			files++
		}
	}
	return false
}
