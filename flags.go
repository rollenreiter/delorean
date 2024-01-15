package main

import "flag"

type cmdflags struct {
	silentFlag bool
	urlFlag    string
	fileFlag   string
}

func GetFlags() *cmdflags {
	var f cmdflags
	flag.BoolVar(&f.silentFlag, "s", false, "Supress all output except for the final link to the archive; useful for scripting")
	flag.StringVar(&f.urlFlag, "u", "", "Declare URLs to archive as a single string or a space-seperated sequence of strings; useful for non-interactive use, scripting and archiving multiple URLs at once")
	flag.StringVar(&f.fileFlag, "f", "", "Declare a file to parse all URLs from; useful for archiving multiple URLs at once")
	flag.Parse()
	return &f
}
