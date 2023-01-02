package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
	"encoding/base64"
)

// read input file into memory, sort, deduplicate, then write to output
// reads from stdin or file with -i file flag -- code needs cleaned up
// writes to stdout or file with -o file flag -- code needs cleaned up
// currently cyclone sort is not as fast as gnu sort, but this provides an alternative that is cross compiled for Linux, Windows & Mac
// writen by cyclone
// v0.3beta-2023.1.1-2300

func versionFunc() {
	fmt.Println("Cyclone Sort")
	fmt.Println("v0.3beta-2023.1.1-2300")
}

func helpFunc() {
	versionFunc()
	fmt.Println("\nSorts and deduplicates by reading entire file into memory.\n")
	fmt.Println("Usage: sort -i input.txt -o output.txt\n")
	fmt.Println("cat input.txt | sort > output.txt\n")
	fmt.Println("Defaults to stdin if -i file is not specified")
	fmt.Println("Defaults to stdout if -o file is not specified\n")
	os.Exit(0)
}

var linesIn int = 0
var deduplicatedLines int = 0

func main() {
	// Parse command-line flags
	inputFileName := flag.String("i", "stdin", "input file (use 'stdin' to read from stdin)")
	outputFileName := flag.String("o", "stdout", "output file (use 'stdout' to write to stdout)")
	version := flag.Bool("version", false, "Program version:")
    cyclone := flag.Bool("cyclone", false, "sort")
    help := flag.Bool("help", false, "Prints help:")
	flag.Parse()

	// run sanity checks for -version, -cyclone & -help
    if *version == true {
        versionFunc()
        os.Exit(0)
    } else if *cyclone == true {
        funcBase64Decode("Q29kZWQgYnkgY3ljbG9uZSA7KQo=")
        os.Exit(0)
    } else if *help == true {
        helpFunc()
    }

	// Launch a goroutine to sort and deduplicate lines
	done := make(chan struct{})
	go func() {
		startTime := time.Now()

		// Open input file
		log.Println("Reading file...")
		var file *os.File
		var err error
		if *inputFileName == "stdin" {
			file = os.Stdin
		} else {
			file, err = os.Open(*inputFileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
		}

		// Create a buffered reader with 10 MB buffer
		reader := bufio.NewReaderSize(file, 10*1024*1024)

		// Read lines of file into a slice
		var lines []string
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
			linesIn++ // update line counter
		}

		// Sort slice
		sort.Strings(lines)

		// Open output file
		var outFile *os.File
		if *outputFileName == "stdout" {
			outFile = os.Stdout
		} else {
			outFile, err = os.Create(*outputFileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer outFile.Close()
		}

		// Create a buffered writer with 10 MB buffer
		writer := bufio.NewWriterSize(outFile,10*1024*1024)

		// Launch a goroutine to write sorted and deduplicated lines to output file -- testing to see if this goroutine helps performance
		var lastLine string
		var deduplicatedLines int
		writeDone := make(chan struct{})
		go func() {
			for _, line := range lines {
				if line != lastLine {
					_, err = writer.WriteString(line + "\n")
					if err != nil {
						fmt.Println(err)
						return
					}
					lastLine = line
					deduplicatedLines++
				}
			}

			// Flush writer's buffer to ensure all data is written to output file
			err = writer.Flush()
			if err != nil {
				fmt.Println(err)
				return
			}
			close(writeDone)
		}()

		// Wait for write goroutine to finish
		<-writeDone

		// Calculate elapsed time
		elapsedTime := time.Since(startTime)

		// Print statistics
		log.Printf("Lines in: %d\n", linesIn)
		log.Printf("Lines out: %d\n", deduplicatedLines)
		log.Printf("Duplicates Removed: %d\n", linesIn-deduplicatedLines)
		log.Printf("Elapsed time: %v\n", elapsedTime)
		log.Printf("Lines Per Second Processed: %.0f\n", float64(linesIn)/elapsedTime.Seconds())

		close(done)
	}()

	// Wait for goroutine to finish
	<-done
}

// base64 decode function used for displaying encoded messages
func funcBase64Decode(line string) {
    str, err := base64.StdEncoding.DecodeString(line)
    if err != nil {
		log.Println("--> Text doesn't appear to be base64 encoded. <--")
        os.Exit(0)
    }
    fmt.Printf("%s\n", str)
}
