
# DeLorean

A simple ~~and blazingly fast~~ CLI for archiving webpages on the [Wayback
Machine](https://web.archive.org) written in *Go*

# Installation

## Building from source

**DeLorean** only uses the Go standard library and does not require any
external dependencies to run or compile.

```bash
git clone --depth=1 https://github.com/rollenreiter/delorean
cd delorean
go build .
```

Although building will not fail, running DeLorean on Windows, macOS, Plan9,
Haiku and the BSDs is **not officially supported.** Proper Windows support
is coming soon.

# Usage

Executing the binary without flags opens an interactive interface.

```bash
# Run DeLorean in interactive mode
$ delorean
Welcome to DeLorean v0.1.0
Type h for help

[0] >
```

Use the `-u` flag to archive a whitespace-separated list of URLs.

```bash
# Archive example.foo, example1.foo and example2.foo
delorean -u "https://example.foo http://example1.foo example2.foo"
```

Use the `-s` flag to suppress all output except for the final output.

```bash
# Archive example.foo, example1.foo and example2.foo
# and redirect all output to citation.txt
delorean -u "https://example.foo https://example1.foo" -s >> citation.txt
```

Use the `-f` flag to pass a file. DeLorean will parse all URLs and archive them.

```bash
# Archive all links from old-citation.txt and redirect
# the output to new-citation.txt
delorean -f old-citation.txt -s >> new-citation.txt
```
