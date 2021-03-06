package main

import(
    "bufio"
    "os"
    "log"
    "fmt"

    "github.com/nathan-fiscaletti/consolesize-go"
)

func main() {
    // Determine the rect of the console window.
    cols, rows := consolesize.GetConsoleSize()
    if rows == 0 || cols == 0 {
        return
    }

    // Subtract one for the input line
    rows = rows - 1

    // Determine if the user passed us a file and if so,
    // handle the file.
    if len(os.Args) >= 2 {
        file, err := os.Open(os.Args[1])
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        handleFile(cols, rows, file)
        return
    }

    // If no file has been handed to us, attempt to use standard input

    // Determine if standard input has any data in it
    file := os.Stdin
    fi, err := file.Stat()
    if err != nil {
        log.Fatal(err)
    }
    size := fi.Size()

    if size > 0 {
        handleFile(cols, rows, os.Stdin)
        return
    }

    // Print the help message if no data was found in standard input
    fmt.Println("peek is a command that will fill your console window")
    fmt.Println("with the contents of the input file until the")
    fmt.Println("console window cannot fit any more text.")
    fmt.Printf("%s", "\n");
    fmt.Println("reading data from standard input is also supported")
    fmt.Printf("%s", "\n");
    fmt.Printf("usage: %s <file>\n", os.Args[0])
}

// handleFile will scan the provided file and print out it's contents
// until the rect described by cols and rows is totally filled.
func handleFile(cols int, rows int, file *os.File) {
    scanner := bufio.NewScanner(file)

    displayedLineCount := 0
    for scanner.Scan() {
        line := scanner.Text()
        displayedLineCount += 1
        if len(line) > cols {
            displayedLineCount += 1
        }

        fmt.Printf("%s\n", line)

        if displayedLineCount >= rows {
            break
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

