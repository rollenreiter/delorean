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
	flag.Usage = func() {
		fmt.Printf("%s|x|%s ", Error, Escape)
		fmt.Printf("Invalid flags or subcommand. Please invoke 'delorean -h' for help.\n")
		os.Exit(1)
	}
	fromfile := flag.NewFlagSet("file", flag.ExitOnError)
	var f CmdArgs
	flag.BoolVar(&f.alphaFlag, "a", false, "Sort URLs in alphabetical order.")
	flag.BoolVar(&f.silentFlag, "s", false, "Supress all output except for the final link to the archive.")
	flag.BoolVar(&f.versionFlag, "v", false, "Print Version and exit.")
	flag.StringVar(&f.fromString, "u", "", "Deprecated: Identical to \"delorean '[...URLS]'\"")
	flag.StringVar(&f.fromString, "f", "", "Deprecated: Identical to 'delorean file'")
	fsilent := fromfile.Bool("s", false, "Supress all output except for the final link to the archive.")
	falpha := fromfile.Bool("a", false, "Sort URLs in alphabetical order.")

	if len(os.Args) == 1 {
		f.fromFile = ""
		f.fromString = ""
	} else {
		switch os.Args[1] {
		case "-u":
			{
				fmt.Printf("%s|!|%s ", Warning, Escape)
				fmt.Printf("The '-u' flag is deprecated and %sWILL%s be removed in a later version. Please use 'delorean [urls] [flags...]' instead.\n\n", Error, Escape)
				f.fromString = os.Args[2]
			}
		case "file", "-f":
			{
				if os.Args[1] == "-f" {
					fmt.Printf("%s|!|%s ", Warning, Escape)
					fmt.Printf("The '-f' flag is deprecated and %sWILL%s be removed in a later version. Please use 'delorean file <filename> [flags...]' instead.\n\n", Error, Escape)
				}
				if len(os.Args) == 2 {
					fmt.Printf("%s|!|%s ", Error, Escape)
					fmt.Println("No file given")
					fmt.Println("USAGE: delorean file [FLAGS]... [FILE]")
					os.Exit(1)
				} else {
					if IsMultifile() {
						fmt.Printf("%s|!|%s ", Warning, Escape)
						fmt.Printf("Using multiple input files is not yet supported. Only the first file will be read.\n\n")
					}
					fromfile.Parse(os.Args[2:])
					f.alphaFlag = *falpha
					f.silentFlag = *fsilent
					f.fromFile = fromfile.Args()[0]
				}
			}
		case "-h":
			{
				fmt.Printf("DeLorean - a CLI for archiving webpages on the Wayback Machine\n")
				fmt.Printf("Usage:\n\n")
				fmt.Printf("\tdelorean [urls] [flags...]\n")
				fmt.Printf("\tdelorean file [flags...] <filename>\n")
				fmt.Printf("\nFlags:\n")
				fmt.Printf("\t-h\tshow this help message\n")
				fmt.Printf("\t-s\tsilence the output except for the final list of URLs\n")
				fmt.Printf("\t-a\tsort the final output alphabetically\n")
				fmt.Printf("\nDeprecated:\n")
				fmt.Printf("\t-u\tlegacy equivalent to 'delorean <urls>'\n")
				fmt.Printf("\t-f\tlegacy equivalent to 'delorean file <file>'\n")
				os.Exit(0)
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
