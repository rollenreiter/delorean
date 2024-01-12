package main

import (
	"fmt"
	"log"
	"strings"
)

func GetUrl(f cmdflags) []string {
	u := []string{""}
	if f.urlFlag == "" {
		fmt.Println("Enter a link to archive:")
		_, err := fmt.Scanln(&u[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		urls := f.urlFlag
		u = strings.Fields(urls)
	}
	return u
}
