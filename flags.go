package main

import (
	"flag"
	"fmt"
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
				if len(os.Args) == 2 {
					fmt.Println("delorean: no file given")
					fmt.Println("USAGE: delorean file [FLAGS]... [FILE]")
					os.Exit(1)
				} else {
					if len(os.Args) > 3 {
						fmt.Printf("%sWARNING:%s ", Warning, Escape)
						fmt.Printf("Using multiple input files is not yet supported. Only the first file will be read.\n\n")
					}
					fromfile.Parse(os.Args[2:])
					f.alphaFlag = *falpha
					f.silentFlag = *fsilent
					f.fromFile = fromfile.Args()[0]
				}
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
