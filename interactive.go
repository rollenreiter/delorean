package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Colors struct {
	bold   string
	reset  string
	red    string
	green  string
	yellow string
	blue   string
	purple string
	cyan   string
	gray   string
	white  string
}

func GetColors() Colors {
	var c Colors
	if runtime.GOOS == "windows" {
		c = Colors{
			"", "", "", "", "", "", "", "", "", "",
		}
	} else {
		c = Colors{
			"\033[1m",
			"\033[0m",
			"\033[31m",
			"\033[32m",
			"\033[33m",
			"\033[34m",
			"\033[35m",
			"\033[36m",
			"\033[37m",
			"\033[97m",
		}
	}
	return c
}

func InterfaceInit(u *urls) {
	version := 1.0
	fmt.Fprintf(os.Stderr, "DeLorean Version %f (interactive mode)\nPress h for help \n", version)
	c := GetColors()

	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(c.blue + "> " + c.reset)
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		switch s {
		case "q":
			os.Exit(1)
		case "a":
			InterfaceAdd(u, &c)
		case "h":
			InterfaceHelp(&c)
		case "l":
			InterfaceList(u, &c)
		}
	}
}

func InterfaceHelp(c *Colors) {
	fmt.Printf("Keybinds:\n")
	fmt.Printf("%sq\t%sQuit interactive interface\n", c.bold, c.reset)
	fmt.Printf("%sl\t%sList all links in the archive list\n", c.bold, c.reset)
	fmt.Printf("%sa\t%sAdd a URL to the archive list\n", c.bold, c.reset)
	fmt.Printf("%sw\t%sArchive all URLs in archive list, then quit\n", c.bold, c.reset)
	fmt.Printf("\n")
}

func InterfaceAdd(u *urls, c *Colors) {
	fmt.Println("Enter a URL to add to the archive list:")
	fmt.Print(c.red + "> " + c.reset)
	var new string
	fmt.Scanln(&new)
	u.tokens = append(u.tokens, new)
	fmt.Printf("Successfully added %s to archive list\n", new)
}

func InterfaceList(u *urls, c *Colors) {
	if len(u.tokens) == 0 {
		fmt.Printf("The archive list is currently %sempty%s\n\n", c.gray, c.reset)
	} else {
		fmt.Printf("These links will be archived on write:\n")
		for i, s := range u.tokens {
			fmt.Printf("(%d) %s\n", i+1, s)
		}
		fmt.Println()
	}
}
