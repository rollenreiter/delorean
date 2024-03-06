package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	Greeter      = fmt.Sprintf("\033[34mWelcome to DeLorean (%s)\033[0m\nType h for help.\n\n", Version)
	Alphabetical = false
)

func InterfaceInit(u *urls) {
	fmt.Print(Greeter)

	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("[" + fmt.Sprint(len(u.tokens)) + "] > ")
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		switch s {
		case "q":
			InterfaceExit(u)
		case "a":
			InterfaceAdd(u)
		case "d":
			InterfaceDelete(u)
		case "h":
			InterfaceHelp()
		case "l":
			InterfaceList(u)
		case "w":
			InterfaceArchive(u)
		case "s":
			InterfaceSort(u)
		default:
			fmt.Printf("\nUnknown command.\nType h for help.\n\n")
		}
	}
}

func InterfaceHelp() {
	fmt.Printf("\nKeybinds:\n")
	fmt.Printf("a - Add a URL to the archive list\n")
	fmt.Printf("l - List all links in the archive list\n")
	fmt.Printf("d - Remove a URL from the archive list\n\n")

	fmt.Printf("s - Toggle alphabetical sorting\n")
	fmt.Printf("w - Archive all URLs in archive list\n")
	fmt.Printf("q - Quit interactive interface\n")
	fmt.Println()
}

func InterfaceAdd(u *urls) {
	fmt.Println("\nEnter the URL to add to the archive list:")
	fmt.Print("[" + fmt.Sprint(len(u.tokens)) + "] + ")
	var new string
	fmt.Scanln(&new)
	if new == "\n" || new == "" || new == " " {
		fmt.Printf("Aborted.\n\n")
	} else {
		newToken := token{
			order:   len(u.tokens),
			content: new,
		}
		u.tokens = append(u.tokens, newToken)
		fmt.Printf("Successfully added to archive list\n\n")
	}
}

func InterfaceDelete(u *urls) {
	fmt.Println("\nWhich URL do you want to remove?:")
	for i, s := range u.tokens {
		fmt.Printf("(%d) %s\n", i+1, s.content)
	}
	fmt.Print("[" + fmt.Sprint(len(u.tokens)) + "] - ")
	var selection string
	fmt.Scanln(&selection)
	deleted, err := strconv.Atoi(selection)
	switch {
	case err != nil:
		fmt.Printf("Please enter a number.\n\n")
	case deleted <= 0 || deleted > len(u.tokens):
		fmt.Printf("This URL is not in the list.\n\n")
	default:
		{
			fmt.Printf("Successfully deleted \"%s\" from archive list.\n\n", u.tokens[deleted-1].content)
			new := make([]token, 0, len(u.tokens)-1)
			for i, s := range u.tokens {
				if i == deleted-1 {
					continue
				}
				new = append(new, s)
			}
			u.tokens = new
		}
	}
}

func InterfaceList(u *urls) {
	if len(u.tokens) == 0 {
		fmt.Printf("The archive list is currently empty.\n\n")
	} else {
		fmt.Printf("These links will be archived on write:\n")
		for i, s := range u.tokens {
			fmt.Printf("(%d) %s\n", i+1, s.content)
		}
		fmt.Println()
	}
}

func InterfaceSort(u *urls) {
	if Alphabetical {
		Alphabetical = false
		Sort(u.tokens, Alphabetical)
		fmt.Printf("Sorted URLs in FIFO order.\n\n")
	} else {
		Alphabetical = true
		Sort(u.tokens, Alphabetical)
		fmt.Printf("Sorted URLs in alphabetical order.\n\n")
	}
}

func InterfaceArchive(u *urls) {
	var wg sync.WaitGroup
	if len(u.tokens) != 0 {
		u.ArchiveInter(&wg)
		u.Finish()
	}
	fmt.Printf("Nothing to archive.\n\n")
}

func (u urls) Finish() {
	fmt.Println("\nSUCCESS! These are the links to the archives:")
	Sort(u.results, Alphabetical)
	for _, s := range u.results {
		fmt.Println(s.content)
	}
	os.Exit(0)
}

func InterfaceExit(u *urls) {
	if len(u.tokens) == 0 {
		fmt.Println("\nGoodbye!")
		os.Exit(0)
	} else {
		fmt.Printf("\nWARNING! The following sites:\n")
		for i, s := range u.tokens {
			fmt.Printf("(%d) %s\n", i+1, s.content)
		}
		fmt.Printf("will NOT be archived. Are you sure you want to quit? [y/N]")
		// TODO: make y/N
		var yn string
		fmt.Scanln(&yn)
		switch yn {
		case "y", "Y", "yes", "Yes", "yEs", "yeS", "yES", "YeS", "YEs":
			fmt.Printf("\nGoodbye!\n")
			os.Exit(0)
		default:
			fmt.Printf("Returning to normal mode.\n\n")
			return
		}
	}
}
