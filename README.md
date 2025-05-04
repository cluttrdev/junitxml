# JUnit XML

This project provides a library and command line tool to process junit xml files.

The supported XML format is based on the [testmoapp/junitxml](https://githubm/com/testmoapp/junitxml)
project which documents the common use of JUnit-style XML files by testing and ci tools.

## Installation

For the library, use:

```sh
go get go.cluttr.dev/junitxml@latest
```

For the command line tool, use:

```sh
go install go.cluttr.dev/junitxml/cmd/junitxml@latest
```


## Usage

### Library

For library usage, see [the package docs](https://pkg.go.dev/go.cluttr.dev/junitxml)
or the following quick-start example:

```go
package main

import (
    "fmt"
    "log"
    "os"

    "go.cluttr.dev/junitxml"
)

func main() {
    file, err := os.Open("junit.xml")
    if err := nil {
        log.Fatal("error opening file: %v", err)
    }
    defer file.Close()

    report, err := junitxml.Parse(file)
    if err != nil {
        log.Fatal("error parsing file: %v", err)
    }

    fmt.Printf("Test execution took %v seconds.\n", report.time)
}
```

### Command line tool

To see what the cli can do, run:

```sh
$ junitxml -h
Process junit xml files.

USAGE
  junitxml [COMMAND] [OPTION]... [ARG]...

COMMANDS
  merge    Merge junit xml files
  version  Show version information
```
