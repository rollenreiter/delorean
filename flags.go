package main

import "flag"

type cmdflags struct {
	silentFlag bool
	urlFlag    string
	fileFlag   string
}

func GetFlags() *cmdflags {
	var f cmdflags
	flag.StringVar(&f.urlFlag, "u", "", "Declare URLs to archive as a single string or a space-seperated sequence of strings")
	flag.StringVar(&f.fileFlag, "f", "", "Declare a file to parse all URLs from")
	flag.BoolVar(&f.silentFlag, "s", false, "Supress all output except for the final link to the archive")
	flag.Parse()
	return &f
}
