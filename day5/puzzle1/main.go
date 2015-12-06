package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// niceFinder returns number of nice strings
func niceFinder(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var niceCount int
	for s.Scan() {
		line := s.Text()
		var vowelCount int
		var twice bool
		var naughty bool
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 'a', 'e', 'i', 'o', 'u':
				vowelCount++
			}
			if i > 0 {
				if !twice && line[i] == line[i-1] {
					twice = true
				}
				switch line[i] {
				case 'b':
					if line[i-1] == 'a' {
						fmt.Printf("%s contains ab\n", line)
						naughty = true
					}
				case 'd':
					if line[i-1] == 'c' {
						fmt.Printf("%s contains cd\n", line)
						naughty = true
					}
				case 'q':
					if line[i-1] == 'p' {
						fmt.Printf("%s contains pq\n", line)
						naughty = true
					}
				case 'y':
					if line[i-1] == 'x' {
						fmt.Printf("%s contains xy\n", line)
						naughty = true
					}
				}
			}
		}
		if vowelCount >= 3 && twice && !naughty {
			niceCount++
		}
	}
	return niceCount
}

func main() {
	fmt.Printf("Nice: %d\n", niceFinder(os.Stdin))
}
