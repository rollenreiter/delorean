package main

import (
	"flag"
	"os"
)

type CmdArgs struct {
	fromString  string
	fromFile    string
	silentFlag  bool
	alphaFlag   bool
	versionFlag bool
}

// Parse CmdArgs from Stdin
func FromStdin() CmdArgs {
	fromfile := flag.NewFlagSet("file", flag.ExitOnError)
	var f CmdArgs
	flag.BoolVar(&f.alphaFlag, "a", false, "Sort URLs in alphabetical order.")
	// flag.StringVar(&f.fromString, "u", "", "Declare URLs to archive as a single string or a space-seperated sequence of strings.")
	// flag.StringVar(&f.fromFile, "f", "", "Declare a file to parse and archive all URLs from.")
	flag.BoolVar(&f.silentFlag, "s", false, "Supress all output except for the final link to the archive.")
	flag.BoolVar(&f.versionFlag, "v", false, "Print Version and exit.")
	fsilent := fromfile.Bool("s", false, "Supress all output except for the final link to the archive.")
	falpha := fromfile.Bool("a", false, "Sort URLs in alphabetical order.")
	if len(os.Args) == 1 {
		f.fromFile = ""
		f.fromString = ""
	} else {
		switch os.Args[1] {
		case "file":
			{
				fromfile.Parse(os.Args[2:])
				f.alphaFlag = *falpha
				f.silentFlag = *fsilent
				f.fromFile = fromfile.Args()[0]
			}
		default:
			{
				f.fromString = os.Args[1]
			}
		}
	}
	flag.Parse()
	return f
}

// Set cmdflags according to the output of a function
func (instance *CmdArgs) SetFlags(result func() CmdArgs) {
	*instance = result()
}
