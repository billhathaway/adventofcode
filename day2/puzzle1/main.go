package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func parse(r io.Reader) (int, error) {
	var l, w, h, sa int
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		count, err := fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		if count != 3 {
			return 0, fmt.Errorf("invalid scanf count of %d", count)
		}
		if err != nil {
			return 0, err
		}
		side1 := 2 * l * w
		side2 := 2 * w * h
		side3 := 2 * h * l
		sa += (side1 + side2 + side3)
		min := side1
		if side2 < min {
			min = side2
		}
		if side3 < min {
			min = side3
		}
		sa += (min / 2)
	}
	return sa, nil
}

func main() {
	sa, err := parse(os.Stdin)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	fmt.Printf("surface area: %d\n", sa)
}
