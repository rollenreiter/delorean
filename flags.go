package main

import "flag"

type cmdflags struct {
	urlFlag    string
	fileFlag   string
	silentFlag bool
	alphaFlag  bool
}

// Parse cmdflags from Stdin
func FromStdin() cmdflags {
	var f cmdflags
	flag.BoolVar(&f.alphaFlag, "a", false, "Return final list of archives in alphabetical order")
	flag.StringVar(&f.urlFlag, "u", "", "Declare URLs to archive as a single string or a space-seperated sequence of strings")
	flag.StringVar(&f.fileFlag, "f", "", "Declare a file to parse and archive all URLs from")
	flag.BoolVar(&f.silentFlag, "s", false, "Supress all output except for the final link to the archive")
	flag.Parse()
	return f
}

// Set cmdflags according to the output of a function
func (instance *cmdflags) SetFlags(result func() cmdflags) {
	*instance = result()
}
