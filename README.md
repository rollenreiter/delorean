# Delorean
A simple and fast CLI for archiving webpages on the [Wayback Machine](https://web.archive.org) written in *Go*

# Installation
## AUR package:
```bash
$ yay -S delorean
```

## Building from source:
**Delorean** only uses the Go standard library and does not require any dependencies to run or compile.
```bash
$ git clone --depth=1 https://github.com/rollenreiter/delorean
$ cd delorean
$ go build .
```

# Usage
Executing the binary results in an interactive interface that can be used to quickly archive a single URL:
```bash
delorean
```

Use the `-u` flag to archive a whitespace-separated list of URLs:
```bash
delorean -u "https://example.foo http://example1.foo example2.foo"
```

Use the `-s` flag to suppress all output except for the links to the archives and error messages:
```bash
delorean -u "https://example.foo https://example1.foo" -s >> citation.txt
```
