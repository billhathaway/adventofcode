package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func parse(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanBytes)
	var floor int
	for s.Scan() {
		val := s.Text()
		switch val {
		case ")":
			floor--
		case "(":
			floor++
		case "\n":
		default:
			return floor, fmt.Errorf("invalid token: [%s]\n", val)
		}
	}
	return floor, nil
}

func main() {
	floor, err := parse(os.Stdin)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Floor: %d\n", floor)
}
