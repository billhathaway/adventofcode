package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func parse(r io.Reader) (int, error) {
	var l, w, h, rl int
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
		side1 := 2 * (l + w)
		side2 := 2 * (w + h)
		side3 := 2 * (h + l)
		min := side1
		if side2 < min {
			min = side2
		}
		if side3 < min {
			min = side3
		}

		sa := l * w * h
		fmt.Printf("line=[%s] min perimeter=%d surface area=%d\n", line, min, sa)
		rl += (min + sa)

	}
	return rl, nil
}

func main() {
	rl, err := parse(os.Stdin)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	fmt.Printf("ribbon length: %d\n", rl)
}
