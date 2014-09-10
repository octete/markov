// markov.go
//
// This is a small program to run the markov chain program
// from the book The Practice of Programming.

package main

import (
	"bufio"
	"fmt"
	"github.com/kr/pretty"
	"os"
	"path/filepath"
	"strings"
)

/* Main variable:
statemap will be a map with state as keys, and a slice
to a set of suffixes. This way, we don't really need a
complicated data structure.
*/
var statemap map[state][]string

// State is each of the two words of a prefix that will initiate
// a markov chain. It will point to at least one suffice.
type state struct {
	pref [2]string
	// suffixes []string // For the suffixes I only need a list/slice
}

// We use prefix to keep the actual prefix we are working on.
// TODO: This should be a slice, rather than an array.
var prefix state

func main() {
	statemap = make(map[state][]string)

	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1>\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	filename := os.Args[1]

	// Read the file and store it in memory.
	// Wouldn't it be better to process as we go?
	lines, err := readLines(filename)
	if err != nil {
		fmt.Printf("Error reading %s", filename)
		os.Exit(2)
	}

	// program
	build(lines)
	pretty.Print(statemap)

}

/*
 */
// build: read input, build prefix map.
func build(lines []string) (err error) {
	for _, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			add(word)
		}
	}
	return nil
}

func lookup() {

}

func addSuffix() {

}

/* We add the word on the existing prefix. */
// add: Add the prefix if it doesn't exist. Otherwise, update table.
// TODO: Do not accept the empty word or ""
func add(word string) {
	// prefix is a global variable, and it has the prefixes for the existing word.
	// update prefix after we've added the word.

	// update prefix
	if statemap[prefix] == nil {
		statemap[prefix] = []string{word}
	} else {
		statemap[prefix] = append(statemap[prefix], word)
	}
	// Change prefix
	prefix.pref[0] = prefix.pref[1]
	prefix.pref[1] = word
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// contains: Search a slice to see if it contains an element
func contains(s []string, e string) bool {
	for _, elem := range s {
		if e == elem {
			return true
		}
	}
	return false
}
