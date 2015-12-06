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
	var position int
	for s.Scan() {
		position++
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
		if floor == -1 {
			return position, nil
		}
	}
	return 0, fmt.Errorf("basement never entered")
}

func main() {
	position, err := parse(os.Stdin)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Position: %d\n", position)
}
