// markov.go
//
// This is a small program to run the markov chain program
// from the book The Practice of Programming.

package markov

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	var nlines, nwords int

	nlines, nwords, _ = build(lines)

	fmt.Printf("The number of lines in %s is %d and has %d words\n",
		filepath.Base(filename), nlines, nwords)
}

// Build the hash table.
func build(lines []string) (nwords, nlines int, err error) {
	for r, line := range lines {
		words := strings.Split(line, " ")
		nwords += len(words)
		nlines = r
	}
	return nwords, nlines, err
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
