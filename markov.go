// markov.go
//
// This is a small program to run the markov chain program
// from the book The Practice of Programming.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	//"string"
)

/* Main variable:
statetab will be a map with state as keys, and a slice
to a set of suffixes. This way, we don't really need a
complicated data structure.
*/
var statetab map[state][]suffix

type suffix struct {
	word string
}

// State is each of the two words of a prefix that will initiate
// a markov chaing. It will point to at least one suffice.
type state struct {
	pref [2]string
	//suffixes []string // For the suffixes I only need a list/slice
}

func main() {
	statetab = make(map[state][]suffix)
	//fmt.Println("Markov chains")
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1>\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
}

// Build the hash table.
func build() {

}

func lookup() {

}

func addSuffix() {

}

func add() {

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
