package main

// boilerplate template used for Advent of Code

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type (
	xy struct {
		x int
		y int
	}
)

var ()

func parse(r io.Reader) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}
	return nil
}

func main() {
	err := parse(os.Stdin)
	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(1)
	}
	fmt.Println("OK")
}
