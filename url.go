package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetUrl(f cmdflags) []string {
	u := []string{""}
	if f.urlFlag == "" {
		if !f.silentFlag {
			fmt.Println("Enter a link to archive:")
		}
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

func GetUrlFile(f cmdflags) []string {
	u := []string{""}
	file, err := os.Open(f.fileFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		u = append(u, s.Text())
	}
	return u
}
