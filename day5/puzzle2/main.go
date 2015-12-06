package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// niceFinder returns number of nice strings
func niceFinder(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var niceCount int
	for s.Scan() {
		line := s.Text()
		var pairFound bool
		var repeats bool
		for i := 1; i < len(line); i++ {
			if pairFound && repeats {
				break
			}
			if !pairFound && strings.Contains(string(line[i+1:]), string(line[i-1:i+1])) {
				pairFound = true
			}
			if i >= 2 && !repeats {
				if line[i] == line[i-2] {
					repeats = true
				}
			}
		}
		fmt.Printf("line=%s pairFound=%v repeats=%v\n", line, pairFound, repeats)
		if pairFound && repeats {
			niceCount++
		}
	}
	return niceCount
}

func main() {
	fmt.Printf("Nice: %d\n", niceFinder(os.Stdin))
}
