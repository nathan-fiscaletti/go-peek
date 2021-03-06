# Peek

Peek is a utility that allows you to fill your current console window with as much data from an input file (or standard input) as will fit. It is conscious of things such as _line wrapping_ and takes into account both the *width* and *height* of your console window.

Peek makes use of [consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go) to accomplish this.

## Installation

```sh
$ git clone https://github.com/nathan-fiscaletti/go-peek.git peek
$ cd peek
$ go build -o /usr/local/bin/peek
```

## Usage Examples

```sh
$ peek some_file.txt
```

```sh
$ cat README.md | peek
```