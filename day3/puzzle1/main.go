package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func parse(r io.Reader) (int, error) {
	houses := make(map[[2]int]int)
	var presents int
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanBytes)
	var row, col int
	var loc [2]int
	houses[loc]++
	presents++
	for s.Scan() {
		val := s.Text()
		switch val {
		case "^":
			row++
			loc[0] = row
			loc[1] = col
			houses[loc]++
			presents++
		case "v":
			row--
			loc[0] = row
			loc[1] = col
			houses[loc]++
			presents++
		case "<":
			col--
			loc[0] = row
			loc[1] = col
			houses[loc]++
			presents++
		case ">":
			col++
			loc[0] = row
			loc[1] = col
			houses[loc]++
			presents++
		case "\n":
		default:
			return 0, fmt.Errorf("invalid token: [%s]\n", val)
		}
	}
	fmt.Printf("presents = %d\n", presents)
	return len(houses), nil
}

func main() {
	houses, err := parse(os.Stdin)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	fmt.Printf("OK %d\n", houses)
}
