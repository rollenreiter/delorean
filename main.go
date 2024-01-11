package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rollenreiter/delorian/parser"
)

func main() {
	url, err := parseUrl()
	if err != nil {
		log.Fatal(err)
	}
	archive(url)
}

func archive(site string) {
	fmt.Println("Archiving...")
	url := fmt.Sprintf("https://web.archive.org/save/%s", site)
	cmd := exec.Command("curl", url)
	html, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("temp.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(html)

	parser.Parse(file)
}

func parseUrl() (string, error) {
	var input []string
	var err error
	input = os.Args
	flag.Parse()
	if input[len(input)-1] == input[0] || input[1] == "" {
		err = fmt.Errorf("please supply a url to archive")
	}
	url := input[len(input)-1]
	return url, err
}
