
# DeLorean

A simple ~~and blazingly fast~~ CLI for archiving webpages on the [Wayback
Machine](https://web.archive.org) written in *Go*

# Installation

## Building from source

**DeLorean** only uses the Go standard library and does not require any
external dependencies to run or compile.
Make sure to append $GOPATH/bin to your $PATH.

```bash
# Install DeLorean from source
$ git clone --depth=1 https://github.com/rollenreiter/delorean
$ cd delorean
$ go install .

# If not done already:
$ export PATH=$PATH:$GOPATH/bin
```

[!WARNING]
> Although building will not fail, running DeLorean on Windows, macOS, Plan9,
> Haiku and the BSDs is **not officially supported.** Proper Windows support
> is coming soon.

# Usage

Executing the binary without arguments opens an interactive interface.

```bash
# Run DeLorean in interactive mode
$ delorean
Welcome to DeLorean v0.1.0
Type h for help

[0] >
```

Pass a whitespace-separated list of URLs as an argument to archive all of them.

```bash
# Archive example.foo, example1.foo and example2.foo
delorean "https://example.foo http://example1.foo example2.foo"
```

Use the `-s` flag to suppress all output except for error messages and the
final output.

```bash
# Archive example.foo, example1.foo and example2.foo
# and redirect all output to citation.txt
delorean "https://example.foo https://example1.foo" -s >> citation.txt
```

Use `delorean file` to pass a file. DeLorean will parse all URLs and
archive them.

```bash
# Archive all links from old-citation.txt and redirect
# the output to new-citation.txt
delorean file old-citation.txt -s >> new-citation.txt
```
